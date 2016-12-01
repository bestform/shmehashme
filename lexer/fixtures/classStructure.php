<?php

use Foo\Bar;

class Foo extends Bar implements BarInterface {

    public function foo($bar, $baz)
    {
        return $bar + $baz;
    }
    static private function bar(){}
    protected function baz(){}
}

print(foo(1, 2));