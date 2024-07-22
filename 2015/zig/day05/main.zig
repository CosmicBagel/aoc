const std = @import("std");

pub fn main() !void {
    std.debug.print("day05\n", .{});

    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = std.fs.File.OpenMode.read_only });
    defer file.close();

    var bufferedReader = std.io.bufferedReader(file.reader());
    var reader = bufferedReader.reader();

    var goodCount: u32 = 0;

    var buffer = [_]u8{0} ** 512;
    while (try reader.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        if (processLine(line)) {
            goodCount += 1;
        }
    }

    std.debug.print("{d}\n", .{goodCount});
}

fn processLine(line: []u8) bool {
    // std.debug.print("{s}\n", .{line});

    // It contains a pair of any two letters that appears at least twice in the
    // string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not
    // like aaa (aa, but it overlaps).
    var duplicatePairFound = false;
    // aaaa
    // axaaaa
    for (0..line.len - 2) |i| {
        const letterPair = line[i .. i + 2];
        const searchStart = i + 2;
        const searchEnd = line.len - 1;
        // std.debug.print("len {d}, i {d}, pair {s}, search start {d}, search end {d}, search slice {s}\n", .{
        //     line.len,
        //     i,
        //     letterPair,
        //     searchStart,
        //     searchEnd,
        //     line[searchStart .. searchEnd + 1],
        // });

        for (searchStart..searchEnd) |o| {
            const other = line[o .. o + 2];
            // std.debug.print("\tcomparing {s} {s}\n", .{ letterPair, other });
            if (std.mem.eql(u8, letterPair, other)) {
                // std.debug.print("\tduplicate found\n", .{});
                duplicatePairFound = true;
                // break :outer;
            }
        }
    }
    if (!duplicatePairFound) {
        // std.debug.print("\tduplicate not found\n", .{});
        return false;
    }

    // It contains at least one letter which repeats with exactly one letter
    // between them, like xyx, abcdefeghi (efe), or even aaa.
    var repeatFound = false;
    for (0..line.len - 2) |i| {
        const slice = line[i .. i + 3];
        if (slice[0] == slice[2]) {
            repeatFound = true;
            // std.debug.print("\tsandwich found\n", .{});
            break;
        }
    }
    if (!repeatFound) {
        // std.debug.print("\tsandwich not found\n", .{});
        return false;
    }

    return true;
}
