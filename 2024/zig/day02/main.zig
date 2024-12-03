const std = @import("std");

const print = std.debug.print;
const mem_size_bytes = 256;
const assumed_list_max = 10;

pub fn main() !void {
    try p2();
}

fn p2() !void {
    const input = @embedFile("test_input.txt");
    print("day02\n", .{});

    var buf = [_]u8{0} ** mem_size_bytes;
    var fba = std.heap.FixedBufferAllocator.init(&buf);
    const allocator = fba.allocator();

    var line_iterator = std.mem.splitScalar(u8, input, '\n');
    var list = try std.ArrayList(i32).initCapacity(allocator, assumed_list_max);

    var safe_lines: u32 = 0;

    while (line_iterator.next()) |line| {
        if (line.len == 0) break;
        var segment_iterator = std.mem.splitScalar(u8, line, ' ');

        while (segment_iterator.next()) |segment| {
            // const first_segment = segment_iterator.next().?;
            const number = try std.fmt.parseInt(i32, segment, 10);
            list.appendAssumeCapacity(number);
        }

        for
        safe_lines += 1;
        print("{any}\n", .{list.items});
        list.clearRetainingCapacity();
    }

    // while (line_iterator.next()) |line| {
    //     if (line.len == 0) break;
    //     var segment_iterator = std.mem.splitScalar(u8, line, ' ');
    //
    //     const first_segment = segment_iterator.next().?;
    //     const first_number = try std.fmt.parseInt(i32, first_segment, 10);
    //     print("{d} ", .{first_number});
    //     var problem_dampener_used = false;
    //
    //     var last_number = first_number; //if (problem_dampener_used) first_number else second_number;
    //     var is_line_safe = true;
    //     var expect_increasing: ?bool = null;
    //     while (segment_iterator.next()) |segment| {
    //         const number = try std.fmt.parseInt(i32, segment, 10);
    //         print("{d} ", .{number});
    //         const diff = @abs(number - last_number);
    //         const is_increasing = number > last_number;
    //         // const expect_increasing = last_increase orelse is_increasing;
    //         if (expect_increasing == null) {
    //             expect_increasing = is_increasing;
    //         }
    //
    //         //its xor, but I can't do xor nicely with bools
    //         // is inc | expect inc | not_match
    //         // 1      | 1          | 0
    //         // 0      | 1          | 1
    //         // 1      | 0          | 1
    //         // 0      | 0          | 0
    //         //
    //         const not_match_pattern = is_increasing and !expect_increasing.? or !is_increasing and expect_increasing.?;
    //         if (diff > 3 or diff == 0 or not_match_pattern) {
    //             if (problem_dampener_used) {
    //                 is_line_safe = false;
    //                 print("not safe\n", .{});
    //                 break;
    //             }
    //
    //             problem_dampener_used = true;
    //             print("PD ", .{});
    //             continue;
    //         }
    //         last_number = number;
    //     }
    //     if (is_line_safe) {
    //         safe_lines += 1;
    //         print("safe\n", .{});
    //     }
    // }

    print("{d}\n", .{safe_lines});
}

fn p1() !void {
    // at lest 1 and at most 3 difference between number
    // no change, or too large of change is unsafe
    // count lines that are have safe number sequences
    // only safe if all numbers are increasing or decreasing (so can't go up AND down in one line)
    const input = @embedFile("input.txt");
    print("day02\n", .{});

    var safe_lines: u32 = 0;
    var line_iterator = std.mem.splitScalar(u8, input, '\n');
    while (line_iterator.next()) |line| {
        if (line.len == 0) break;
        var segment_iterator = std.mem.splitScalar(u8, line, ' ');

        const first_segment = segment_iterator.next().?;
        const second_segment = segment_iterator.next().?;
        const first_number = try std.fmt.parseInt(i32, first_segment, 10);
        const second_number = try std.fmt.parseInt(i32, second_segment, 10);

        const inital_diff = @abs(first_number - second_number);
        if (inital_diff > 3 or inital_diff == 0) {
            continue; //not safe
        }
        const expect_increasing = second_number > first_number;

        var last_number = second_number;
        var is_line_safe = true;
        while (segment_iterator.next()) |segment| {
            const number = try std.fmt.parseInt(i32, segment, 10);
            const diff = @abs(number - last_number);
            const is_increasing = number > last_number;
            //its xor, but I can't do xor nicely with bools
            // is inc | expect inc | not_match
            // 1      | 1          | 0
            // 0      | 1          | 1
            // 1      | 0          | 1
            // 0      | 0          | 0
            //
            const not_match_pattern = is_increasing and !expect_increasing or !is_increasing and expect_increasing;
            if (diff > 3 or diff == 0 or not_match_pattern) {
                is_line_safe = false;
                break;
            }
            last_number = number;
        }
        if (is_line_safe) safe_lines += 1;
    }

    print("{d}\n", .{safe_lines});
}
