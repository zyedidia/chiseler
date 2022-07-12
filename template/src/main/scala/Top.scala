package main

import chisel3._

class Top() extends Module {}

object Top extends App {
  (new chisel3.stage.ChiselStage).emitVerilog(new Top(), Array("--target-dir", "generated"))
}
