PollController:
  lda allowPolling
  beq PollControllerDone
    lda #$01
    sta CONTROLLER1
    lda #$00
    sta CONTROLLER1
    ldx #$08
PollControllerLoop:
  lda CONTROLLER1
  lsr
  rol controller1
  dex
  bne PollControllerLoop
    jsr HandleController1
PollControllerDone:
  jmp PollController

HandleController1:
  lda readingData
  beq HandleControllerCheckSyncByte
    jmp HandleController1IOByte
HandleControllerCheckSyncByte:
  lda controller1
  cmp #(SYNC_SATELLITE)
  beq HandleController1SetSyncByte
    cmp #(SYNC_HOME_SUN_MOON)
    beq HandleController1SetSyncByte
      rts
HandleController1SetSyncByte:
  sta dataType
  lda #$01
  sta readingData
  lda #$00
  sta readingDataOffset
  rts
HandleController1IOByte:
  ldx readingDataOffset
  lda dataType
  cmp #(SYNC_SATELLITE)
  beq HandleController1IOSatellite
    jmp HandleController1IONotSatellite
HandleController1IOSatellite:
  lda SatelliteIO, X
  tay
  lda controller1
  sta ControllerRAMData, y
  inx
  stx readingDataOffset
  cpx #40 ; SatelliteIO length
  bcc HandleController1IOSatelliteDone
    lda #$00
    sta readingData
    jsr DrawTopBarData
HandleController1IOSatelliteDone:
  rts
HandleController1IONotSatellite:
  cmp #(SYNC_HOME_SUN_MOON)
  beq HandleController1IOHomeSunMoon
    jmp HandleController1IONotHomeSunMoon
HandleController1IOHomeSunMoon:
  lda HomeSunMoonIO, X
  tay
  lda controller1
  sta ControllerRAMData, y
  inx
  stx readingDataOffset
  cpx #8 ; HomeSunMoonIO length
  bcc HandleController1IOHomeSunMoonDone
    lda #$00
    sta readingData
HandleController1IOHomeSunMoonDone:
  rts
HandleController1IONotHomeSunMoon:
  rts

DrawTopBarData:
  lda #$01
  sta disableDraw
  jsr WaitFrame

  ; TODO
  ldy #$00
  sty drawBufferOffset
  ; TODO

  jsr DrawDate
  jsr DrawTime
  jsr DrawLatitude
  jsr DrawLongitude
  ldy drawBufferOffset
  lda #$8F
  sta (drawBuffer), Y
  iny
  sty drawBufferOffset
  
  jsr LatitudeToTmp
  jsr LatitudeToYPosition
  jsr LongitudeToTmp
  jsr LongitudeToXScroll

  ; Because sun and moon are relational to the satellite,
  ; we execute this immediately after dealing with satellite data
  jsr SetHome
  jsr SetSunMoon
  jsr DrawInViewSprite

  lda #$01
  sta forceDraw
  jsr WaitFrame

  ldy #$00
  sty drawBufferOffset

  jsr DrawShading
  jsr DrawAltitude
  jsr DrawNoradID

  ldy drawBufferOffset
  lda #$8F
  sta (drawBuffer), Y
  iny
  sty drawBufferOffset

  lda #$01
  sta forceDraw
  jsr WaitFrame

  lda #$00
  sta disableDraw
DrawTopBarDataDone:
  rts


DrawShading:
  lda homeSunMoonStatus
  and #%00000010
  bne DrawShadingSunEnabled
    ldx #$00
    jmp DrawShadingContinue
DrawShadingSunEnabled:
  lda homeSunMoonStatus
  and #%01000000
  eor #%01000000
  lsr
  ror
  ror
  ror
  tax
DrawShadingContinue:
  ldy drawBufferOffset
  lda #$3F
  sta (drawBuffer), Y
  iny
  lda #$10
  sta (drawBuffer), Y
  iny
DrawShadingLoop:
  lda ShadowPalette, X
  sta (drawBuffer), Y
  iny
  inx
  txa
  and #%00000011
  bne DrawShadingLoop
    lda #$8E
    sta (drawBuffer), Y
    iny
    sty drawBufferOffset
    rts

DrawAltitude:
  lda #$00
  sta altitudeBlank
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$76
  sta (drawBuffer), Y
  iny
  ldx #$00
DrawAltitudeLoop:
  lda altitude, X
  bne DrawAltitudeLoopNotBlank
    dec altitudeBlank
    inc altitudeBlank
    bne DrawAltitudeLoopNotBlank
      dec altitudeBlank
      lda #$23
DrawAltitudeLoopNotBlank:
  inc altitudeBlank
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  inx
  cpx #$08
  bne DrawAltitudeLoop
    lda #$8E
    sta (drawBuffer), Y
    iny
    sty drawBufferOffset
    rts

DrawNoradID:
  lda #$00
  sta noradIDBlank
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$B5
  sta (drawBuffer), Y
  iny
  ldx #$00
DrawNoradIDLoop:
  lda noradID, X
  bne DrawNoradIDLoopNotBlank
    dec noradIDBlank
    inc noradIDBlank
    bne DrawNoradIDLoopNotBlank
      dec noradIDBlank
      lda #$23
DrawNoradIDLoopNotBlank:
  inc noradIDBlank
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  inx
  cpx #$09
  bne DrawNoradIDLoop
    lda #$8E
    sta (drawBuffer), Y
    iny
    sty drawBufferOffset
    rts

DrawDate:
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$49
  sta (drawBuffer), Y
  iny
  ldx month
DrawDateCheckMonth:
  beq DrawDateInvalidMonth
    cpx #13
    bcs DrawDateInvalidMonth
      dex
      stx tmp
      txa
      asl
      clc
      adc tmp
      tax
      jmp DrawDateContinue
DrawDateInvalidMonth:
  ldx #$00
DrawDateContinue:
  lda Months, X
  sta (drawBuffer), Y
  iny
  inx
  lda Months, X
  sta (drawBuffer), Y
  iny
  inx
  lda Months, X
  sta (drawBuffer), Y
  iny
  lda #$04
  sta (drawBuffer), Y
  iny
  lda day
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (day + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$04
  sta (drawBuffer), Y
  iny
  lda year
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (year + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (year + 2)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (year + 3)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$8E
  sta (drawBuffer), Y
  iny
  sty drawBufferOffset
  rts

DrawTime:
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$69
  sta (drawBuffer), Y
  iny
  lda hour
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (hour + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$E0
  sta (drawBuffer), Y
  iny
  lda minute
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (minute + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$E0
  sta (drawBuffer), Y
  iny
  lda second
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (second + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda satStatus
  and #%00000111
  asl
  asl
  tax
DrawUTCWordLoop:
  lda TimeLabel, X
  sta (drawBuffer), Y
  iny
  inx
  txa
  and #$03
  cmp #$00
  bne DrawUTCWordLoop
    lda #$8E
    sta (drawBuffer), Y
    iny
    sty drawBufferOffset
    rts

DrawLatitude:
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$89
  sta (drawBuffer), Y
  iny
  lda satLat
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (satLat + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$DF
  sta (drawBuffer), Y
  iny
  lda (satLat + 2)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (satLat + 3)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$04
  sta (drawBuffer), Y
  iny
  lda satStatus
  and #%10000000
  asl
  rol
  tax
  lda LatitudeLetter, X
  sta (drawBuffer), Y
  iny
  lda #$8E
  sta (drawBuffer), Y
  iny
  sty drawBufferOffset
  rts

DrawLongitude:
  ldy drawBufferOffset
  lda #$20
  sta (drawBuffer), Y
  iny
  lda #$A8
  sta (drawBuffer), Y
  iny
  lda satLon
  beq DrawLongitudeNotHundred
    clc
    adc #$E1
    jmp DrawLongitudeContinue
DrawLongitudeNotHundred:
  lda #$04
DrawLongitudeContinue:
  sta (drawBuffer), Y
  iny
  lda (satLon + 1)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (satLon + 2)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$DF
  sta (drawBuffer), Y
  iny
  lda (satLon + 3)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda (satLon + 4)
  clc
  adc #$E1
  sta (drawBuffer), Y
  iny
  lda #$04
  sta (drawBuffer), Y
  iny
  lda satStatus
  and #%01000000
  asl
  rol
  rol
  tax
  lda LongitudeLetter, X
  sta (drawBuffer), Y
  iny
  lda #$8E
  sta (drawBuffer), Y
  iny
  sty drawBufferOffset
  rts

SatelliteIO:
  .byte #<(day), #<(day + 1)
  .byte #<(month)
  .byte #<(year), #<(year + 1), #<(year + 2), #<(year + 3)
  .byte #<(hour), #<(hour + 1)
  .byte #<(minute), #<(minute + 1)
  .byte #<(second), #<(second + 1)
  .byte #<(altitude), #<(altitude + 1), #<(altitude + 2), #<(altitude + 3), #<(altitude + 4), #<(altitude + 5), #<(altitude + 6), #<(altitude + 7)
  .byte #<(noradID), #<(noradID + 1), #<(noradID + 2), #<(noradID + 3), #<(noradID + 4), #<(noradID + 5), #<(noradID + 6), #<(noradID + 7), #<(noradID + 8)
  .byte #<(satLat), #<(satLat + 1), #<(satLat + 2), #<(satLat + 3)
  .byte #<(satLon), #<(satLon + 1), #<(satLon + 2), #<(satLon + 3), #<(satLon + 4)
  .byte #<(satStatus)
SatelliteIODone:


HomeSunMoonIO:
  .byte #<(homeLat), #<(homeLon)
  .byte #<(sunLat), #<(sunLon)
  .byte #<(moonLat), #<(moonLon)
  .byte #<(moonPhase)
  .byte #<(homeSunMoonStatus)
HomeSunMoonIODone:
