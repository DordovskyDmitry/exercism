pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        match self.health == 0 {
            true if self.level >= 10 => Some(Self { health: 100, mana: Some(100), ..*self }),
            true => Some(Self { health: 100, ..*self }),
            false => None,
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            Some(v) if v > mana_cost => {
                self.mana = Some(v - mana_cost);
                mana_cost * 2
            }
            Some(v) if v == mana_cost => {
                self.mana = None;
                mana_cost * 2
            }
            Some(_) => 0,
            None if self.health > mana_cost => {
                self.health -= mana_cost;
                0
            },
            None => {
                self.health = 0;
                0
            },
        }
    }
}
