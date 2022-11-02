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

.enum $0000
Variables:
  nmi                 dsb 1
  screen              dsb 1
  addr                dsb 2
  addrEnd             dsb 2
  ppuAddr             dsb 2
  xscroll             dsb 1
  nametable           dsb 1
  patterns            dsb 1
  tmp                 dsb 2

  skipFrame           dsb 1 
  disableDraw         dsb 1
.ende

.enum $0300
DrawBufferVariables:
  drawBufferOffset      dsb 1
  drawBufferOffsetTmp   dsb 1
  drawBuffer            dsb 200
  forceDraw             dsb 1
.ende

.enum $0400
TopBarVariables:
  satelliteDataWriting  dsb 1
  satelliteDataOffset   dsb 2
.ende

.enum $0500
ControllerRAMData:
  allowPolling          dsb 1
  controller1           dsb 1
  readingData           dsb 1
  readingDataOffset     dsb 1
  dataType              dsb 1

  satStatus             dsb 1
  day                   dsb 2
  month                 dsb 1
  year                  dsb 4
  hour                  dsb 2
  minute                dsb 2
  second                dsb 2
  altitude              dsb 8
  noradID               dsb 9
  satLat                dsb 4
  satLon                dsb 5


  homeLat               dsb 1
  homeLon               dsb 1
  sunLat                dsb 1
  sunLon                dsb 1
  moonLat               dsb 1
  moonLon               dsb 1
  moonPhase             dsb 1
  homeSunMoonStatus     dsb 1

  altitudeBlank         dsb 1
  noradIDBlank          dsb 1
.ende

.enum $0600
  crewmateDataWriting   dsb 1
  crewmateDataOffset    dsb 2
  crewmateBuffer        dsb 508
.ende
