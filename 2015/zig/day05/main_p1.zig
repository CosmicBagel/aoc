const std = @import("std");

var ptBuffer = [_]u8{0} ** 4096;
var ptAlloc = std.heap.FixedBufferAllocator.init(&ptBuffer);
var prefixTree: *PrefixTree = undefined;

pub fn main() !void {
    std.debug.print("day05\n", .{});

    var pt = try PrefixTree.create(ptAlloc.allocator());
    prefixTree = &pt;
    // bad strings ab, cd, pq, or xy
    try prefixTree.insert("ab");
    try prefixTree.insert("cd");
    try prefixTree.insert("pq");
    try prefixTree.insert("xy");

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
    // It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    var vowelCount: u8 = 0;
    for (line) |char| {
        switch (char) {
            'a', 'e', 'i', 'o', 'u' => {
                vowelCount += 1;
            },
            else => {},
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

    for (0..line.len - 2) |i| {
        const chars = line[i .. i + 2];
        if (prefixTree.search(chars)) {
            return false;
        }
    }

    return true;
}

const PTNode = struct {
    //26 letters in (ascii supported) english alphabet
    children: [26]?*PTNode,
    terminal: bool,
};

const PrefixTree = struct {
    root: *PTNode,
    alloc: std.mem.Allocator,

    fn create(allocator: std.mem.Allocator) !PrefixTree {
        const n = try allocator.create(PTNode);
        n.* = std.mem.zeroInit(PTNode, .{});
        return PrefixTree{ .root = n, .alloc = allocator };
    }

    fn insert(self: PrefixTree, string: []const u8) !void {
        var n = self.root;
        for (string) |c| {
            // check if uppercase, then subtract that
            const i = if (c < 0x61) c - 0x41 else c - 0x61;
            const next = try self.alloc.create(PTNode);
            next.* = std.mem.zeroInit(PTNode, .{});
            n.children[i] = next;
            n = next;
        }
        n.terminal = true;
    }

    fn search(self: PrefixTree, string: []const u8) bool {
        var n = self.root;
        for (string) |c| {
            // check if uppercase, then subtract that
            const i = if (c < 0x61) c - 0x41 else c - 0x61;
            if (n.children[i] != null) {
                n = n.children[i].?;
            } else {
                return false;
            }
        }

        return n.terminal;
    }
};
