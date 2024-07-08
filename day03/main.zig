const std = @import("std");

pub fn main() !void {
    std.debug.print("day03", .{});

    var f = try std.fs.cwd().openFile("day03-input-test.txt", .{});
    defer f.close();

    var bufferedReader = std.io.bufferedReader(f.reader());
    const reader = bufferedReader.reader();

    // hashset for visited coordinates
    // bump 'visted' count when new entry is made

    var x: i32 = 0;
    var y: i32 = 0;
    var visitedCount: u32 = 0;

    const Coord = struct { x: i32, y: i32 };
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    var visitedMap = std.AutoHashMap(Coord, void).init(
        gpa.allocator(),
    );
    defer visitedMap.deinit();

    // count the starting point
    try visitedMap.put(Coord{ .x = 0, .y = 0 }, {});
    visitedCount += 1;

    while (true) {
        const byte = reader.readByte() catch |err| switch (err) {
            error.EndOfStream => break,
            else => |e| return e,
        };

        switch (byte) {
            '>' => x += 1,
            '<' => x -= 1,
            '^' => y += 1,
            'v' => y -= 1,
            '\n' => {},
            else => {
                std.debug.print("surprise character '{c}', 0x{x}\n", .{byte, byte});
                return error.Oops;
            },
        }

        const coord = Coord{ .x = x, .y = y };
        const contains = visitedMap.contains(coord);
        if (!contains) {
            visitedCount += 1;
            try visitedMap.put(coord, {});
        }
    }

    std.debug.print("visitedCount: {}\n", .{visitedCount});
}
