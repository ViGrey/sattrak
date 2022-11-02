LatitudeToTmp:
  lda #$00
  ldx satLat
LatitudeToTmpTensLoop:
  cpx #$00
  beq LatitudeToTmpTensLoopDone
    clc
    adc #$0A
    dex
    jmp LatitudeToTmpTensLoop
LatitudeToTmpTensLoopDone:
  clc
  adc (satLat + 1)
  cmp #91
  bcc LatitudeToTmpDone
    lda #$00
LatitudeToTmpDone:
  sta (tmp + 1)
  rts

LatitudeToYPosition:
  lda satStatus
  and #%10000000
  beq LatitudeToYPositionSouth
    lda #90
    sec
    sbc (tmp + 1)
    jmp LatitudeToYPositionContinue
LatitudeToYPositionSouth:
  lda (tmp + 1)
  clc
  adc #90
LatitudeToYPositionContinue:
  tay
  lda (Latitudes), Y
  sta $208
  clc
  adc #$03
  sta $204
  sta $20C
  rts

LongitudeToTmp:
  lda #$00
  ldx satLon
LongitudeToTmpHundredsLoop:
  beq LongitudeToTmpHundredsLoopDone
    clc
    adc #100
LongitudeToTmpHundredsLoopDone:
  ldx (satLon + 1)
LongitudeToTmpTensLoop:
  cpx #$00
  beq LongitudeToTmpTensLoopDone
    clc
    adc #$0A
    dex
    jmp LongitudeToTmpTensLoop
LongitudeToTmpTensLoopDone:
  clc
  adc (satLon + 2)
  cmp #180
  bcc LongitudeToTmpDone
    lda #180
LongitudeToTmpDone:
  sta (tmp + 1)
  rts

LongitudeToXScroll:
  lda satStatus
  and #%01000000
  bne LongitudeToXScrollEast
    lda #$00
    sta nametable
    lda #180
    sec
    sbc (tmp + 1)
    jmp LongitudeToXScrollContinue
LongitudeToXScrollEast:
  lda #$01
  sta nametable
  lda (tmp + 1)
LongitudeToXScrollContinue:
  tay
  lda (Longitudes), Y
  sec
  sbc #$80
  sta xscroll
  bcs LongitudeToXScrollDone
    lda nametable
    eor #%00000001
    sta nametable
LongitudeToXScrollDone:
  lda (tmp + 1)
  sta satLon
  rts












;;;;;;;;;;
;;
;; MAP Screen Draw Code
;;
;;;;;;;;;;

DrawMap:
  lda #$00
  sta allowPolling
  jsr BlankScreen
  lda #$01
  sta screen
  lda #%00001000
  sta patterns
  jsr ClearPPURAM
  jsr SetMapScreenPalettes
  jsr DrawMapScreen
  jsr ClearSprites
  jsr DrawSpriteZero
  jsr DrawSatSprites
  jsr DrawSunSprite
  jsr DrawMoonSprite
  jsr DrawHomeSprite
  jsr DrawTopBar
  jsr InitializeDrawBuffer
  jsr LatitudeToTmp
  jsr LatitudeToYPosition
  jsr LongitudeToTmp
  jsr LongitudeToXScroll
  lda #$00
  sta disableDraw
  jsr ResetScroll
  lda #$01
  sta forceDraw
  jsr WaitFrame
  jsr DrawTopBarData
  rts

SetMapScreenPalettes:
  lda PPU_STATUS
  lda #$3F
  sta PPU_ADDR
  lda #$00
  sta PPU_ADDR
  lda #<(Palettes)
  sta addr
  lda #>(Palettes)
  sta (addr + 1)
  ldy #$00
SetMapScreenPalettesLoop:
  lda (addr), Y
  sta PPU_DATA
  iny
  cpy #$20
  bne SetMapScreenPalettesLoop
    rts

DrawMapScreen:
  lda PPU_STATUS
  lda #$20
  sta PPU_ADDR
  lda #$00
  sta PPU_ADDR
  lda #<(West)
  sta addr
  lda #>(West)
  sta (addr + 1)
  ldy #$00
  ldx #$04
DrawMapScreenWest:
  lda (addr), Y
  sta PPU_DATA
  iny
  bne DrawMapScreenWest
    inc (addr + 1)
    dex
    bne DrawMapScreenWest
      lda PPU_STATUS
      lda #$24
      sta PPU_ADDR
      lda #$00
      sta PPU_ADDR
      lda #<(East)
      sta addr
      lda #>(East)
      sta (addr + 1)
      ldx #$04
DrawMapScreenEast:
  lda (addr), Y
  sta PPU_DATA
  iny
  bne DrawMapScreenEast
    inc (addr + 1)
    dex
    bne DrawMapScreenEast
      rts

DrawSunSprite:
  lda #$FE
  sta $21C
  sta $220

  lda #$8F
  sta $21D
  sta $221

  lda #$02
  sta $21E
  lda #$42
  sta $222

  rts

DrawMoonSprite:
  lda #$FE
  sta $214
  sta $218

  lda #$97
  sta $215
  sta $219

  lda #$01
  sta $216
  lda #$41
  sta $21A

  rts

DrawHomeSprite:
  lda #$FE
  sta $224
  lda #$A1
  sta $225
  lda #$01
  sta $226

  rts


;;;;;;;;;;
;;
;; SET HOME SPRITE
;;
;;;;;;;;;;

SetHome:
  lda homeSunMoonStatus
  and #%00000001
  bne SetHomeHomeEnabled
    lda #$FE
    sta $224
    rts
SetHomeHomeEnabled:
  lda #180
  sec
  sbc homeLat
  tax
  lda Latitudes, X
  sta $224
HomeLonHandling:
  lda #$01
  sta (tmp + 1)
  ldy homeLon
  lda satStatus
  and #%00001000
  bne HomeLonEast
    dec (tmp + 1)
    lda #180
    sec
    sbc homeLon
    tay
HomeLonEast:

  lda (Longitudes), Y
  sec
  sbc xscroll
  sta $227

  bcs HomeLonGTEXScroll
    lda (tmp + 1)
    eor nametable
    and #$01
    bne HomeLonContinue
      lda #$FE
      sta $224
      jmp HomeLonContinue

HomeLonGTEXScroll:
  lda (tmp + 1)
  eor nametable
  and #$01
  beq HomeLonContinue
    lda #$FE
    sta $224
HomeLonContinue:
  lda $227
  sec
  sbc #$04
  sta $227
  bcs HomeLonGTE8
    lda $224
    cmp #$FE
    beq HomeLonSet
      jmp HomeLonClear
HomeLonGTE8:
  lda $224
  cmp #$FE
  bne HomeLonSet
HomeLonClear:
  lda #$FE
  sta $224
HomeLonSet:
  rts



SetSunMoon:
SetSunMoonSunCheck:
  lda homeSunMoonStatus
  and #%00000010
  bne SetSunMoonSunEnabled
    lda #$FE
    sta $21C
    sta $220
    jmp SetSunMoonMoonCheck
SetSunMoonSunEnabled:
  lda #180
  sec
  sbc sunLat
  tax
  lda Latitudes, X
  sta $21C
  sta $220
SetSunMoonMoonCheck:
  lda homeSunMoonStatus
  and #%00000100
  bne SetSunMoonMoonEnabled
    lda #$FE
    sta $214
    sta $218
    jmp SunLonHandling
SetSunMoonMoonEnabled:
  lda #180
  sec
  sbc moonLat
  tax
  lda Latitudes, X
  sta $214
  sta $218

  lda moonPhase
  and #%00001111
  asl
  tax
  lda MoonPhases, X
  ora #%00000001
  sta $215
  lda MoonPhases, X
  lsr
  lda #$04
  ror
  lsr
  sta $216
  inx
  lda MoonPhases, X
  ora #%00000001
  sta $219
  lda MoonPhases, X
  lsr
  lda #$04
  ror
  lsr
  sta $21A

SunLonHandling:
  lda #$01
  sta (tmp + 1)
  ldy sunLon
  lda homeSunMoonStatus
  and #%00010000
  bne LonSignsSunEast
    dec (tmp + 1)
    lda #180
    sec
    sbc sunLon
    tay
LonSignsSunEast:
  lda (Longitudes), Y
  sec
  sbc xscroll
  sta $223
  bcs SunLonGTEXScroll
    lda (tmp + 1)
    eor nametable
    and #$01
    bne SunLonContinue
      lda #$FE
      sta $220
      jmp SunLonContinue
SunLonGTEXScroll:
  lda (tmp + 1)
  eor nametable
  and #$01
  beq SunLonContinue
    lda #$FE
    sta $220
SunLonContinue:

  lda $223
  sec
  sbc #$08
  sta $21F
  bcs SunLonGTE8
    ;; Underflow happening
    lda $220
    cmp #$FE
    beq SunLonLeftContinue
      lda #$FE
      sta $21C
      jmp SunLonLeftContinue
SunLonGTE8:
  lda $220
  cmp #$FE
  bne SunLonLeftContinue
    lda #$FE
    sta $21C
SunLonLeftContinue:
  




MoonLonHandling:
  lda #$01
  sta (tmp + 1)
  ldy moonLon
  lda homeSunMoonStatus
  and #%00100000
  bne LonSignsMoonEast
    dec (tmp + 1)
    lda #180
    sec
    sbc moonLon
    tay
LonSignsMoonEast:
  lda (Longitudes), Y
  sec
  sbc xscroll
  sta $21B
  bcs MoonLonGTEXScroll
    lda (tmp + 1)
    eor nametable
    and #$01
    bne MoonLonContinue
      lda #$FE
      sta $218
      jmp MoonLonContinue
MoonLonGTEXScroll:
  lda (tmp + 1)
  eor nametable
  and #$01
  beq MoonLonContinue
    lda #$FE
    sta $218
MoonLonContinue:

  lda $21B
  sec
  sbc #$08
  sta $217
  bcs MoonLonGTE8
    ;; Underflow happening
    lda $218
    cmp #$FE
    beq MoonLonLeftContinue
      lda #$FE
      sta $214
      jmp MoonLonLeftContinue
MoonLonGTE8:
  lda $218
  cmp #$FE
  bne MoonLonLeftContinue
    lda #$FE
    sta $214
MoonLonLeftContinue:
  rts

DrawInViewSprite:
  lda homeSunMoonStatus
  and #%00000001
  beq DrawInViewSpriteNotInView
    lda homeSunMoonStatus
    and #%10000000
    beq DrawInViewSpriteNotInView
      lda $20C
      sec
      sbc #$09
      jmp DrawInViewSpriteContinue
DrawInViewSpriteNotInView:
  lda #$FE
DrawInViewSpriteContinue:
  sta $210
  lda #$87
  sta $211
  lda #$01
  sta $212
  lda $20F
  sta $213
  rts
