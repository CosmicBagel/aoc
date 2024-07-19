const std = @import("std");

pub fn main() !void {
    std.debug.print("day03", .{});

    var f = try std.fs.cwd().openFile("day03-input-test.txt", .{});
    defer f.close();

    var bufferedReader = std.io.bufferedReader(f.reader());
    const reader = bufferedReader.reader();

    // hashset for visited coordinates
    // bump 'visted' count when new entry is made

    const Coord = struct { x: i32, y: i32 };

    var santa = Coord{ .x = 0, .y = 0 };
    var robot = Coord{ .x = 0, .y = 0 };
    var flipFlop = true;
    var visitedCount: u32 = 0;

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

        var coord = if (flipFlop) &santa else &robot;
        flipFlop = !flipFlop;

        switch (byte) {
            '>' => coord.x += 1,
            '<' => coord.x -= 1,
            '^' => coord.y += 1,
            'v' => coord.y -= 1,
            '\n' => {},
            else => {
                std.debug.print("surprise character '{c}', 0x{x}\n", .{ byte, byte });
                return error.Oops;
            },
        }

        const contains = visitedMap.contains(coord.*);
        if (!contains) {
            visitedCount += 1;
            try visitedMap.put(coord.*, {});
        }
    }

    std.debug.print("visitedCount: {}\n", .{visitedCount});
}
