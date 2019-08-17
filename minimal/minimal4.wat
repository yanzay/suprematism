(module
  (func $i (import "imports" "log") (param i32))
  (func (export "print42")
    i32.const 42
    call $i))