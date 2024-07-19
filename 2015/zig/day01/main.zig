const std = @import("std");

pub fn main() !void {
    var f = try std.fs.cwd().openFile("day01-input.txt", .{});
    defer f.close();

    const bufferSize = 512;
    var buf = [_]u8{0} ** bufferSize;

    var d = doritos{};
    while (true) {
        const bytesRead = try f.read(&buf);
        if (bytesRead > 0) {
            d.step(buf[0..bytesRead]);
        } else {
            break;
        }
    }

    std.debug.print("final floor {d}\n", .{d.floor});
}

const doritos = struct {
    floor: i32 = 0,
    charCount: i32 = 0,
    gate: bool = true,

    fn step(self: *doritos, chars: []u8) void {
        for (chars) |c| {
            self.charCount += 1;
            switch (c) {
                '(' => {
                    self.floor += 1;
                },
                ')' => {
                    self.floor -= 1;
                },
                else => {},
            }

            if (self.gate and self.floor == -1) {
                self.gate = false;
                std.debug.print("first -1 floor char at {d}\n", .{self.charCount});
            }
        }
    }
};
