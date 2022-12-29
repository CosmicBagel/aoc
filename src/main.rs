use std::{collections::HashSet, slice::Windows};

use crate::helper::read_lines;

mod helper;

fn main() {
    p1();
}

const SAMPLE_DATA: &str = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k";

enum NodeData {
    Directory,
    File(u32),
}

struct Node {
    parent: Option<usize>,
    children: Vec<usize>,
    node_data: NodeData,
    node_name: String,
}

struct Tree {
    nodes: Vec<Node>,
}

impl Tree {
    //new fn (creates tree with empty root node always at index 0)
    //insert node (parent: usize, data, name)
    //traversal (recursively go through tree, until all nodes are exhausted)
    //no moving nodes, no removing nodes

    fn new() -> Tree {
        let mut t = Tree { nodes: Vec::new() };
        t.nodes.push(Node {
            parent: None,
            children: Vec::new(),
            node_data: NodeData::Directory,
            node_name: String::from("root"),
        });

        t
    }

    fn add_child_to_node(&mut self, data: NodeData, node_name: String, parent: usize) -> usize {
        // ensure child is unique in name among children
        // (no files or dirs can have the same name, and be siblings)
        let mut exists = false;
        for n in &self.nodes[parent].children {
            if self.nodes[*n].node_name == node_name {
                exists = true;
            }
        }
        if exists {
            panic!(
                "sibling with same name already exists (adding: {})",
                node_name
            );
        }

        // node name is unique, can now add

        let next_index = self.nodes.len();

        let n = Node {
            children: Vec::new(),
            parent: Some(parent),
            node_data: data,
            node_name: node_name,
        };

        self.nodes.push(n);
        self.nodes[parent].children.push(next_index);

        next_index
    }

    fn traverse_all(&self) {
        self.traverse_from_node(0, 0);
    }

    fn traverse_from_node(&self, node: usize, depth: u32) {
        // print this node
        let n = &self.nodes[node];
        let indent = "----".repeat(depth as usize);
        print!("{}", indent);
        match n.node_data {
            NodeData::Directory => print!(" dir - {}", n.node_name),
            NodeData::File(size) => print!(" file - {} - {}", n.node_name, size),
        };
        print!("\n");

        // call on all child nodes
        for c in &n.children {
            self.traverse_from_node(*c, depth + 1);
        }
    }
}

fn p1() {
    let mut tree = Tree::new();

    let mut current_parent: usize = 0;

    for line in SAMPLE_DATA.lines() {
        let words: Vec<&str> = line.split(" ").collect();
        let first_word_chars: Vec<char> = words[0].chars().collect();
        let second_word_chars: Vec<char> = words[1].chars().collect();
        let first_word = words[0];
        let second_word = words[1];

        match first_word_chars[0] {
            '$' => {
                //command
                match second_word_chars[0] {
                    'c' => {
                        // change directory (third word is directory name or .. to go up)
                        let dir_name: Vec<char> = words[2].chars().collect();
                        if dir_name.len() == 1 && dir_name[0] == '/' {
                            current_parent = 0; // 0 is root
                            continue;
                        } else if dir_name.len() == 2 && dir_name[0] == '.' && dir_name[1] == '.' {
                            match tree.nodes[current_parent].parent {
                                Some(val) => {
                                    current_parent = val;
                                }
                                None => {}
                            }
                            continue;
                        }

                        // dir name is a child (should already be added by previous list command)
                        let dir_name = words[2];
                        for n in &tree.nodes[current_parent].children {
                            if tree.nodes[*n].node_name == dir_name {
                                current_parent = *n;
                                continue;
                            }
                        }
                    }
                    'l' => {
                        // list, doesn't really matter, anything not a cd is a listing
                    }
                    _ => println!("unknown command: {}", second_word_chars[0]),
                }
            }
            'd' => {
                // directory listing

                //check if already added as child

                // doesn't exist, add as child
                tree.add_child_to_node(
                    NodeData::Directory,
                    String::from(second_word),
                    current_parent,
                );
            }
            '0'..='9' => {
                // file
                // filename
                let file_size: u32 = first_word.parse().unwrap();
                tree.add_child_to_node(
                    NodeData::File(file_size),
                    String::from(second_word),
                    current_parent,
                );
            }

            _ => (),
        }
    }
    tree.traverse_all();
}
