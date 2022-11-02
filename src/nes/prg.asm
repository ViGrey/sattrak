; Copyright (C) 2020, Vi Grey
; All rights reserved.
;
; Redistribution and use in source and binary forms, with or without
; modification, are permitted provided that the following conditions
; are met:
;
; 1. Redistributions of source code must retain the above copyright
;    notice, this list of conditions and the following disclaimer.
; 2. Redistributions in binary form must reproduce the above copyright
;    notice, this list of conditions and the following disclaimer in the
;    documentation and/or other materials provided with the distribution.
;
; THIS SOFTWARE IS PROVIDED BY AUTHOR AND CONTRIBUTORS ``AS IS'' AND
; ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
; IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
; ARE DISCLAIMED. IN NO EVENT SHALL AUTHOR OR CONTRIBUTORS BE LIABLE
; FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
; DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
; OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
; HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
; LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
; OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
; SUCH DAMAGE.

.include "controller.asm"

RESET:
  sei
  cld
  ldx #$40
  stx APU_FRAME_COUNTER
  ldx #$FF
  txs
  inx
  lda #%00000110
  sta PPU_MASK
  lda #$00
  sta PPU_CTRL
  stx $4010
  ldy #$00

InitialVWait:
  lda PPU_STATUS
  bpl InitialVWait
InitialVWait2:
  lda PPU_STATUS
  bpl InitialVWait2

InitializeRAM:
  ldx #$00
InitializeRAMLoop:
  lda #$00
  sta $0000, x
  sta $0100, x
  sta $0300, x
  sta $0400, x
  sta $0500, x
  sta $0600, x
  sta $0700, x
  lda #$FE
  sta $0200, x
  inx
  bne InitializeRAMLoop
    jsr ResetScroll
    jsr DrawMap

jmp PollController

NMI:
  php
  pha
  txa
  pha
  tya
  pha
  inc nmi
  lda #$00
  sta PPU_OAM_ADDR
  lda #$02
  sta OAM_DMA

;  lda skipFrame
;  beq NMINotSkipFrame
;    lda #$00
;    sta skipFrame
;    jmp NMIDone
NMINotSkipFrame:

  jsr HandleDraw
  lda screen
  beq NMIHandleScreenNotMap

NMISprite0ClearWait:
  bit PPU_STATUS
  bvs NMISprite0ClearWait
NMISprite0HitWait:
  bit PPU_STATUS
  bvc NMISprite0HitWait
    lda #%10100000
    ora nametable
    ora patterns
    sta PPU_CTRL
    lda xscroll
    sta PPU_SCROLL

NMIHandleScreenNotMap:
  jsr Update
NMIDone:
  jsr EndDrawBuffer
  pla
  tay
  pla
  tax
  pla
  plp
  rti

Update:
  rts


HandleDraw:
  lda forceDraw
  bne HandleDrawContinue
    lda disableDraw
    bne HandleDrawDone
HandleDrawContinue:
  lda PPU_STATUS
  jsr Draw
  jsr DrawPreviousFrame
  lda #$01
  sta allowPolling
  jsr ResetScroll
HandleDrawDone:
  rts

WaitFrame:
  lda nmi
WaitFrameLoop:
  cmp nmi
  beq WaitFrameLoop
    rts

.include "decompress.asm"
.include "draw.asm"
.include "include.asm"
.include "screen.asm"
.include "map.asm"
