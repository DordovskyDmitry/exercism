pub fn divmod(dividend: i16, divisor: i16) -> (i16, i16) {
    let c = dividend / divisor;
    (c, dividend - c * divisor)
}

pub fn evens<T>(iter: impl Iterator<Item = T>) -> impl Iterator<Item = T> {
    iter.enumerate().filter(|(i, _)| i % 2 == 0).map(|(_, v)| v)
}

pub struct Position(pub i16, pub i16);
impl Position {
    pub fn manhattan(&self) -> i16 {
        let Position(x,y) = self;
        x.abs() + y.abs()
    }
}
