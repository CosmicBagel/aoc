const std = @import("std");

pub fn main() !void {
    var f = try std.fs.cwd().openFile("day02-input.txt", .{});
    defer f.close();

    var bufferedReader = std.io.bufferedReader(f.reader());
    const reader = bufferedReader.reader();

    const bufferSize = 512;
    var buf = [_]u8{0} ** bufferSize;
    var sum: i32 = 0;
    var ribbonLength: i32 = 0;
    while (try reader.readUntilDelimiterOrEof(
        &buf,
        '\n',
    )) |line| {
        var iter = std.mem.splitScalar(u8, line, 'x');
        const l = try std.fmt.parseInt(i32, iter.next().?, 10);
        const w = try std.fmt.parseInt(i32, iter.next().?, 10);
        const h = try std.fmt.parseInt(i32, iter.next().?, 10);

        const top = l * w;
        const front = w * h;
        const right = h * l;

        const minFace = @min(right, @min(top, front));

        const boxSurfaceArea = 2 * top + 2 * front + 2 * right;
        sum += boxSurfaceArea + minFace;

        const topPerimiter = 2 * (l + w);
        const frontPerimiter = 2 * (w + h);
        const rightPerimiter = 2 * (h + l);

        const minPerimiter = @min(rightPerimiter, @min(topPerimiter, frontPerimiter));
        const volume = l * h * w;
        ribbonLength += minPerimiter + volume;
    }

    std.debug.print("{d}\n", .{sum});
    std.debug.print("{d}\n", .{ribbonLength});
}
