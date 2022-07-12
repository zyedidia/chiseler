package test

import chisel3._
import chiseltest._
import org.scalatest.flatspec.AnyFlatSpec

class {{ .Top }}Test extends AnyFlatSpec with ChiselScalatestTester {
  "{{ .Top }} test" should "pass" in {
    test(new {{ .Package }}.{{ .Top }}) { dut =>
      // ...
    }
  }
}
