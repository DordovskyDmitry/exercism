// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut magazine_words = HashMap::new();
    for word in magazine.iter() {
        match magazine_words.get(word) {
            Some(n) => magazine_words.insert(word, n + 1),
            None => magazine_words.insert(word, 1),
        };
    }

    for note_word in note.iter() {
        match magazine_words.get(note_word) {
            Some(1) => magazine_words.remove(note_word),
            Some(n) => magazine_words.insert(note_word, n - 1),
            None => return false,
        };
    }
    true
}
