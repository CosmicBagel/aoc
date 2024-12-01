const std = @import("std");

const print = std.debug.print;
const assumed_input_size = 1000;

pub fn main() !void {
    const input = @embedFile("input.txt");
    //pull out left and right columns into separate lists
    //sort the lists
    //run through them in parallel finding abs diff of each row
    //return sum of diffs

    var mem_arr = [_]u8{0} ** 16384;
    var fba = std.heap.FixedBufferAllocator.init(&mem_arr);
    const allocator = fba.allocator();

    var left_column = try std.ArrayList(i32).initCapacity(allocator, assumed_input_size);
    var right_column = try std.ArrayList(i32).initCapacity(allocator, assumed_input_size);

    var line_iterator = std.mem.splitScalar(u8, input, '\n');
    while (line_iterator.next()) |line| {
        if (line.len < 1) continue;
        var segment_iterator = std.mem.splitScalar(u8, line, ' ');
        const left = try std.fmt.parseInt(i32, segment_iterator.next().?, 10);

        while (segment_iterator.peek()) |segment| {
            if (segment.len > 0) {
                break;
            }
            _ = segment_iterator.next();
        }

        const right = try std.fmt.parseInt(i32, segment_iterator.next().?, 10);
        left_column.appendAssumeCapacity(left);
        right_column.appendAssumeCapacity(right);
    }

    std.mem.sort(i32, left_column.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, right_column.items, {}, comptime std.sort.asc(i32));

    var sum: u32 = 0;
    for (left_column.items, right_column.items) |left, right| {
        // const result, _ = @subWithOverflow(left, right);
        sum += @abs(left - right);
    }
    print("{d}\n", .{sum});
}
