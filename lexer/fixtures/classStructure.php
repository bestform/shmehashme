<?php

use Foo\Bar;

class Foo {

    public function foo($bar, $baz)
    {
        return $bar + $baz;
    }
    private function bar(){}
    protected function baz(){}
}

print(foo(1, 2));