#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack = vec![];

    for elem in inputs.iter() {
        match elem {
            CalculatorInput::Value(n) => { stack.push(*n) },
            _ => {
                match (stack.pop(), stack.pop()) {
                    (Some(v1), Some(v2)) => {
                        match elem {
                            CalculatorInput::Add => stack.push(v2 + v1),
                            CalculatorInput::Subtract => stack.push(v2 - v1),
                            CalculatorInput::Multiply => stack.push(v2 * v1),
                            CalculatorInput::Divide => stack.push(v2 / v1),
                            _ => return None,
                        }
                    },
                    _ => return None,
                }
            }
        }
    }

    if stack.len() != 1 {
        None
    } else {
        stack.pop()
    }
}
