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

ResetScroll:
  lda #$00
  sta PPU_SCROLL
  sta PPU_SCROLL
  jsr EnableNMI
  rts

Draw:
  lda #%00011000
  sta PPU_MASK
  rts

DisableNMI:
  lda #$00
  sta PPU_CTRL
  rts

EnableNMI:
  lda #%10100000
  ora patterns
  sta PPU_CTRL
  rts

Blank:
  lda #%00000000
  sta PPU_MASK
  jsr DisableNMI
  rts

ClearPPURAM:
  lda PPU_STATUS
  lda #$20
  sta PPU_ADDR
  lda #$00
  sta PPU_ADDR
  ldy #$10
  ldx #$00
  txa
ClearPPURAMLoop:
  sta PPU_DATA
  dex
  bne ClearPPURAMLoop
    ldx #$00
    dey
    bne ClearPPURAMLoop
      rts

DrawPreviousFrame:
  ldy #$00
DrawPreviousFrameLoop:
  lda PPU_STATUS
  lda (drawBuffer), Y
  iny
  cmp #$8F
  beq DrawPreviousFrameDone
    cmp #$8E
    beq DrawPreviousFrameLoop
      sta PPU_ADDR
      lda (drawBuffer), Y
      iny
      sta PPU_ADDR
DrawPreviousFrameLoopContentLoop:
  lda (drawBuffer), Y
  iny
  cmp #$8E
  beq DrawPreviousFrameLoop
    ; Not #$8E
    cmp #$8F
    beq DrawPreviousFrameDone
      ; Not #$8F
      sta PPU_DATA
      jmp DrawPreviousFrameLoopContentLoop
DrawPreviousFrameDone:
  jsr InitializeDrawBuffer
  rts

InitializeDrawBuffer:
  ldy #$00
  sty drawBufferOffset
  lda #$8F
  sta (drawBuffer), Y
  rts

EndDrawBuffer:
  lda disableDraw
  bne EndDrawBufferDone
    ldy drawBufferOffset
    lda #$8F
    sta (drawBuffer), Y
    iny
    sty drawBufferOffset
EndDrawBufferDone:
  rts

ClearSprites:
  ldy #$00
  lda #$FE
ClearSpritesLoop:
  sta $200, Y
  iny
  bne ClearSpritesLoop
    rts

DrawSpriteZero:
  lda #$27
  sta $200
  lda #$D1
  sta $201
  lda #$01
  sta #$202
  lda #08
  sta $203
  rts

DrawSatSprites:
  lda #$82
  sta $204
  lda #$81
  sta $205
  lda #$7F
  sta $208
  lda #$83
  sta $209
  lda #$82
  sta $20C
  lda #$85
  sta $20D
  lda #$78
  sta $207
  lda #$80
  sta $20B
  lda #$88
  sta $20F
  lda #$00
  sta $206
  sta $20A
  sta $20E
  rts

DrawISSSprites:
  lda #$6F
  sta $204
  lda #$B0
  sta $205
  lda #$00
  sta $206
  lda #$78
  sta $207
  lda #$6F
  sta $208
  lda #$B1
  sta $209
  lda #$00
  sta $20A
  lda #$80
  sta $20B
  lda #$6F
  sta $20C
  lda #$B2
  sta $20D
  lda #$00
  sta $20E
  lda #$88
  sta $20F
  lda #$77
  sta $210
  lda #$C0
  sta $211
  lda #$00
  sta $212
  lda #$78
  sta $213
  lda #$77
  sta $214
  lda #$C1
  sta $215
  lda #$00
  sta $216
  lda #$80
  sta $217
  lda #$77
  sta $218
  lda #$C2
  sta $219
  lda #$00
  sta $21A
  lda #$88
  sta $21B
  rts

DrawTopBar:
  lda #<(TopBarTextDone)
  sta addrEnd
  lda #>(TopBarTextDone)
  sta (addrEnd + 1)
  lda #<(TopBarText)
  sta addr
  lda #>(TopBarText)
  sta (addr + 1)
  jsr DecompressAddr
  rts

DrawStatusIcons:
  jsr DrawTAStm32StatusIcon
  jsr DrawInternetStatusIcon
  rts

;UpdateStatusIcons:
;  lda screen
;  bne UpdateStatusIconsMap
;    lda statusTAStm32
;    beq UpdateStatusIconsTAStm32Crewmate
;      jsr ResetYTAStm32StatusIcon
;      jmp UpdateStatusIconsCrewmateContinue
;UpdateStatusIconsTAStm32Crewmate:
;  jsr HandleTAStm32StatusIconCrewmate
;UpdateStatusIconsCrewmateContinue:
;  lda statusInternet
;  beq UpdateStatusIconsInternetCrewmate
;    jsr ResetYInternetStatusIcon
;    rts
;UpdateStatusIconsInternetCrewmate:
;  jsr HandleInternetStatusIconCrewmate
;  rts
;UpdateStatusIconsMap:
;  lda statusTAStm32
;  beq UpdateStatusIconsTAStm32Map
;    jsr ResetYTAStm32StatusIcon
;    jmp UpdateStatusIconsMapContinue
;UpdateStatusIconsTAStm32Map:
;  jsr HandleTAStm32StatusIconMap
;UpdateStatusIconsMapContinue:
;  lda statusInternet
;  beq UpdateStatusIconsInternetMap
;    jsr ResetYInternetStatusIcon
;    rts
;UpdateStatusIconsInternetMap:
;  jsr HandleInternetStatusIconMap
;  rts

DrawTAStm32StatusIcon:
  lda #$83
  sta $231
  lda #$84
  sta $235
  lda #$93
  sta $239
  lda #$94
  sta $23D
  lda #$02
  sta $236
  sta $232
  sta $23A
  sta $23E
  rts

DrawInternetStatusIcon:
  lda #$85
  sta $241
  lda #$86
  sta $245
  lda #$95
  sta $249
  lda #$96
  sta $24D
  lda #$02
  sta $246
  sta $242
  sta $24A
  sta $24E
  rts

HandleTAStm32StatusIconMap:
  lda #$0E
  sta $230
  sta $234
  lda #$16
  sta $238
  sta $23C
  lda #$C8
  sta $233
  sta $23B
  lda #$D0
  sta $237
  sta $23F
  rts

HandleInternetStatusIconMap:
  lda #$0E
  sta $240
  sta $244
  lda #$16
  sta $248
  sta $24C
  lda #$E0
  sta $243
  sta $24B
  lda #$E8
  sta $247
  sta $24F
  rts

HandleTAStm32StatusIconCrewmate:
  lda #$23
  sta $230
  sta $234
  lda #$2B
  sta $238
  sta $23C
  lda #$1D
  sta $233
  sta $23B
  lda #$24
  sta $237
  sta $23F
  rts

HandleInternetStatusIconCrewmate:
  lda #$23
  sta $240
  sta $244
  lda #$2B
  sta $248
  sta $24C
  lda #$D4
  sta $243
  sta $24B
  lda #$DC
  sta $247
  sta $24F
  rts

ResetYTAStm32StatusIcon:
  lda #$FE
  sta $230
  sta $234
  sta $238
  sta $23C
  rts

ResetYInternetStatusIcon:
  lda #$FE
  sta $240
  sta $244
  sta $248
  sta $24C
  rts


BlankScreen:
  lda #$01
  sta disableDraw
  lda nmi
BlankScreenWaitFrame:
  cmp nmi
  beq BlankScreenWaitFrame
    jsr SetBlankPalette
    jsr Blank
    rts

