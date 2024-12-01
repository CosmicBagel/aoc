const std = @import("std");

const print = std.debug.print;

const grid_size = 1000 * 1000;
// const grid_size = 3 * 3;

pub fn main() !void {
    std.debug.print("day 06\n", .{});

    // array of 1 million elements (1000x1000)
    // 2d array in flat array using row * row_size + column
    // column major (x,y) coords
    // 0,0 bottom left corner of grid

    // instructions: turn on, turn off, toggle
    // each instruction has a corrdinate pair indicating a rectangle
    //  (bottom-left, top-right), the coordinates are *inclusive*
    // eg
    // turn on 0,0 through 999,999
    // toggle 0,0 through 999,0
    // turn off 499,499 through 500,500
    // ---
    // 998996 will be left on

    // count lights lit up at end of instruction list

    // ingest text line by line from txt file (input-test.txt)
    // we can load the whole file in
    const input_text = @embedFile("input_test.txt");
    // print("{any}\n", .{input_text});
    var grid = [_]bool{false} ** grid_size;
    var lines_iterator = std.mem.splitAny(u8, input_text, "\n");
    printGridCount(&grid);
    while (lines_iterator.next()) |line| {
        if (line.len == 0) break;
        print("{s}\n", .{line});
        const parsed_line = try parseLine(line);
        // print("{any}\n", .{parsed_line});
        gridChange(&grid, parsed_line);
        print("\n", .{});
        printGridCount(&grid);
    }
    printGridCount(&grid);
}

fn gridCount(grid: *[grid_size]bool) u32 {
    var count: u32 = 0;
    for (grid) |cell| {
        if (cell) count += 1;
    }
    return count;
}

fn printGridCount(grid: *[grid_size]bool) void {
    print("light count: {d}\n", .{gridCount(grid)});
}

fn gridChange(grid: *[grid_size]bool, command: Line) void {
    const rect = command.rect;
    print("{any}: {d},{d};{d},{d}\n", .{
        command.command,
        rect.top,
        rect.left,
        rect.bottom,
        rect.right,
    });

    const height = rect.bottom - rect.top;
    const width = rect.right - rect.left;

    for (0..height) |row| {
        for (0..width) |column| {
            const index = (rect.top + row) * 1000 + rect.left + column;
            switch (command.command) {
                .On => grid.*[index] = true,
                .Off => grid.*[index] = false,
                .Toggle => grid.*[index] = !grid.*[index],
            }
        }
    }
}

const LineCommand = enum {
    On,
    Off,
    Toggle,
};

const Rect = struct {
    bottom: u32 = 0,
    left: u32 = 0,
    top: u32 = 0,
    right: u32 = 0,
};

const Line = struct {
    command: LineCommand = LineCommand.On,
    rect: Rect = Rect{},
};

fn parseLine(line: []const u8) !Line {
    // "turn on 0,0 through 999,999"
    //parse command
    //parse range
    //
    var parsed_line = Line{};

    //split on spaces
    //check second char of first word u vs o
    var words_iterator = std.mem.splitAny(u8, line, " ");
    const first_word = words_iterator.next();
    const second_word = words_iterator.next();

    const is_toggle = (first_word.?[1] == 'o');
    if (is_toggle) {
        parsed_line.command = LineCommand.Toggle;
    } else {
        const onOff = second_word.?[1] == 'n';
        parsed_line.command = if (onOff) LineCommand.On else LineCommand.Off;
    }

    // parse range
    const start_str = if (is_toggle) second_word else words_iterator.next();
    //skip "through"
    _ = words_iterator.next();
    const end_str = words_iterator.next();

    var start_coord_iterator = std.mem.splitAny(u8, start_str.?, ",");
    parsed_line.rect.left = try std.fmt.parseInt(u32, start_coord_iterator.next().?, 10);
    parsed_line.rect.top = try std.fmt.parseInt(u32, start_coord_iterator.next().?, 10);
    var end_coord_iterator = std.mem.splitAny(u8, end_str.?, ",");
    parsed_line.rect.right = try std.fmt.parseInt(u32, end_coord_iterator.next().?, 10);
    parsed_line.rect.bottom = try std.fmt.parseInt(u32, end_coord_iterator.next().?, 10);

    return parsed_line;
}

test "turn on 3x3 in top left of grid" {
    const rect = Rect{
        .top = 0,
        .left = 0,
        .bottom = 3,
        .right = 3,
    };
    const command = LineCommand.On;
    const line = Line{
        .rect = rect,
        .command = command,
    };
    var grid = [_]bool{false} ** grid_size;
    gridChange(&grid, line);
    try std.testing.expect(gridCount(&grid) == 9);
}

test "turn on 3x3 in bottom right of grid" {
    const rect = Rect{
        .top = 999 - 3,
        .left = 999 - 3,
        .bottom = 999,
        .right = 999,
    };
    const command = LineCommand.On;
    const line = Line{
        .rect = rect,
        .command = command,
    };
    var grid = [_]bool{false} ** grid_size;
    gridChange(&grid, line);
    try std.testing.expect(gridCount(&grid) == 9);
}

test "turn on all lights" {
    const rect = Rect{
        .top = 0,
        .left = 0,
        .bottom = 999,
        .right = 999,
    };
    const command = LineCommand.On;
    const line = Line{
        .rect = rect,
        .command = command,
    };
    var grid = [_]bool{false} ** grid_size;
    gridChange(&grid, line);
    const count = gridCount(&grid);
    print("light count {d}\n", .{count});
    try std.testing.expect(count == 1_000_000);
}
