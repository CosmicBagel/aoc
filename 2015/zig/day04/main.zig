const std = @import("std");
const Md5 = std.crypto.hash.Md5;

pub fn main() !void {
    // abcdef  -> 609043
    // pqrstuv -> 1048970

    std.debug.print("day04\n", .{});

    // const keyBase = "abcdef";
    // const keyBase = "pqrstuv";
    const keyBase = "yzbqklnj";

    for (0..std.math.maxInt(u64)) |count| {
        var hashOut: [16]u8 = undefined;
        var key: [32]u8 = undefined;
        const formattedString = try std.fmt.bufPrint(&key, "{s}{d}", .{ keyBase, count });
        Md5.hash(formattedString, &hashOut, .{});

        var value: u24 = @bitCast(hashOut[0..3].*);
        value &= 0xffffff;

        if (value == 0) {
            std.debug.print("found hash: 0x{x}\t{d}\n", .{ hashOut, count });
            break;
        }
    }
}
