SetPalette:
  lda #<(Palettes)
  sta addr
  lda #>(Palettes)
  sta (addr + 1)
  lda PPU_STATUS
  lda #$3F
  sta PPU_ADDR
  lda #$00
  sta PPU_ADDR
  ldy #$00
SetPaletteLoop:
  lda (addr), Y
  sta PPU_DATA
  iny
  cpy #$20
  bne SetPaletteLoop
    rts

SetBlankPalette:
  lda PPU_STATUS
  lda #$3F
  sta PPU_ADDR
  lda #$00
  sta PPU_ADDR
  ldx #$20
  lda #$0F
SetBlankPaletteLoop:
  sta PPU_DATA
  dex
  bne SetBlankPaletteLoop
    rts
