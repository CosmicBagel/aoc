const std = @import("std");

pub fn main() !void {
    std.debug.print("day05\n", .{});

    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = std.fs.File.OpenMode.read_only });
    defer file.close();

    var bufferedReader = std.io.bufferedReader(file.reader());
    var reader = bufferedReader.reader();

    var good_count: u32 = 0;

    var buffer = [_]u8{0} ** 512;
    while (try reader.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        if (processLine(line)) {
            good_count += 1;
        }
    }

    std.debug.print("{d}\n", .{good_count});
}

fn processLine(line: []u8) bool {
    // It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    var vowelCount: u8 = 0;
    for (line) |char| {
        switch (char) {
            'a', 'e', 'i', 'o', 'u' => {
                vowelCount += 1;
            },
        }
    }
    if (vowelCount < 3) {
        return false;
    }

    // It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    var lastChar: u8 = line[0];
    var doubleFound = false;
    for (line[1..]) |char| {
        if (char == lastChar) {
            doubleFound = true;
            break;
        }
        lastChar = char;
    }
    if (!doubleFound) {
        return false;
    }

    // It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.



    return true;
}
