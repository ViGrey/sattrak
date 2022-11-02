package main

import (
 // "math"
)
/*

import (
	"github.com/joshuaferrara/go-satellite"

	"math"
	"time"
)

var (
	//TABLE 47.A
	//MEEUS PAGES 339 & 340
	sumLRArray = [][]float64{
		[]float64{0, 0, 1, 0, 6288774, -20905355}, //1
		[]float64{2, 0, -1, 0, 1274027, -3699111}, //2
		[]float64{2, 0, 0, 0, 658314, -2955968},   //3
		[]float64{0, 0, 2, 0, 213618, -569925},    //4
		[]float64{0, 1, 0, 0, -185116, 48888},     //5
		[]float64{0, 0, 0, 2, -114332, -3149},     //6
		[]float64{2, 0, -2, 0, 58793, 246158},     //7
		[]float64{2, -1, -1, 0, 57066, -152138},   //8
		[]float64{2, 0, 1, 0, 53322, -170733},     //9
		[]float64{2, -1, 0, 0, 45758, -204586},    //10
		[]float64{0, 1, -1, 0, -40923, -129620},   //11
		[]float64{1, 0, 0, 0, -34720, 108743},     //12
		[]float64{0, 1, 1, 0, -30383, 104755},     //13
		[]float64{2, 0, 0, -2, 15327, 10321},      //14
		[]float64{0, 0, 1, 2, -12528, 0},          //15
		[]float64{0, 0, 1, -2, 10980, 79661},      //16
		[]float64{4, 0, -1, 0, 10675, -34782},     //17
		[]float64{0, 0, 3, 0, 10034, -23210},      //18
		[]float64{4, 0, -2, 0, 8548, -21636},      //19
		[]float64{2, 1, -1, 0, -7888, 24208},      //20
		[]float64{2, 1, 0, 0, -6766, 30824},       //21
		[]float64{1, 0, -1, 0, -5163, -8379},      //22
		[]float64{1, 1, 0, 0, 4987, -16675},       //23
		[]float64{2, -1, 1, 0, 4036, -12831},      //24
		[]float64{2, 0, 2, 0, 3994, -10445},       //25
		[]float64{4, 0, 0, 0, 3861, -11650},       //26
		[]float64{2, 0, -3, 0, 3665, 14403},       //27
		[]float64{0, 1, -2, 0, -2689, -7003},      //28
		[]float64{2, 0, -1, 2, -2602, 0},          //29
		[]float64{2, -1, -2, 0, 2390, 10056},      //30
		[]float64{1, 0, 1, 0, -2348, 6322},        //31
		[]float64{2, -2, 0, 0, 2236, -9884},       //32
		[]float64{0, 1, 2, 0, -2120, 5751},        //33
		[]float64{0, 2, 0, 0, -2069, 0},           //34
		[]float64{2, -2, -1, 0, 2048, -4950},      //35
		[]float64{2, 0, 1, -2, -1773, 4130},       //36
		[]float64{2, 0, 0, 2, -1595, 0},           //37
		[]float64{4, -1, -1, 0, 1215, -3958},      //38
		[]float64{0, 0, 2, 2, -1110, 0},           //39
		[]float64{3, 0, -1, 0, -892, 3258},        //40
		[]float64{2, 1, 1, 0, -810, 2616},         //41
		[]float64{4, -1, -2, 0, 759, -1897},       //42
		[]float64{0, 2, -1, 0, -713, -2117},       //43
		[]float64{2, 2, -1, 0, -700, 2354},        //44
		[]float64{2, 1, -2, 0, 691, 0},            //45
		[]float64{2, -1, 0, -2, 596, 0},           //46
		[]float64{4, 0, 1, 0, 549, -1423},         //47
		[]float64{0, 0, 4, 0, 537, -1117},         //48
		[]float64{4, -1, 0, 0, 520, -1571},        //49
		[]float64{1, 0, -2, 0, -487, -1739},       //50
		[]float64{2, 1, 0, -2, -399, 0},           //51
		[]float64{0, 0, 2, -2, -381, -4421},       //52
		[]float64{1, 1, 1, 0, 351, 0},             //53
		[]float64{3, 0, -2, 0, -340, 0},           //54
		[]float64{4, 0, -3, 0, 330, 0},            //55
		[]float64{2, -1, 2, 0, 327, 0},            //56
		[]float64{0, 2, 1, 0, -323, 1165},         //57
		[]float64{1, 1, -1, 0, 299, 0},            //58
		[]float64{2, 0, 3, 0, 294, 0},             //59
		[]float64{2, 0, -1, -2, 0, 8752}}          //60

	//TABLE 47.B
	//MEEUS PAGES 341
	//Eb Periodic terms for the latitude of the Moon
	sumBArray = [][]float64{
		[]float64{0, 0, 0, 1, 5128122}, //1
		[]float64{0, 0, 1, 1, 280602},  //2
		[]float64{0, 0, 1, -1, 277693}, //3
		[]float64{2, 0, 0, -1, 173237}, //4
		[]float64{2, 0, -1, 1, 55413},  //5
		[]float64{2, 0, -1, -1, 46271}, //6
		[]float64{2, 0, 0, 1, 32573},   //7
		[]float64{0, 0, 2, 1, 17198},   //8
		[]float64{2, 0, 1, -1, 9266},   //9
		[]float64{0, 0, 2, -1, 8822},   //10
		[]float64{2, -1, 0, -1, 8216},  //11
		[]float64{2, 0, -2, -1, 4324},  //12
		[]float64{2, 0, 1, 1, 4200},    //13
		[]float64{2, 1, 0, -1, -3359},  //14
		[]float64{2, -1, -1, 1, 2463},  //15
		[]float64{2, -1, 0, 1, 2211},   //16
		[]float64{2, -1, -1, -1, 2065}, //17
		[]float64{0, 1, -1, -1, -1870}, //18
		[]float64{4, 0, -1, -1, 1828},  //19
		[]float64{0, 1, 0, 1, -1794},   //20
		[]float64{0, 0, 0, 3, -1749},   //21
		[]float64{0, 1, -1, 1, -1565},  //22
		[]float64{1, 0, 0, 1, -1491},   //23
		[]float64{0, 1, 1, 1, -1475},   //24
		[]float64{0, 1, 1, -1, -1410},  //25
		[]float64{0, 1, 0, -1, -1344},  //26
		[]float64{1, 0, 0, -1, -1335},  //27
		[]float64{0, 0, 3, 1, 1107},    //28
		[]float64{4, 0, 0, -1, 1021},   //29
		[]float64{4, 0, -1, 1, 833},    //30
		[]float64{0, 0, 1, -3, 777},    //31
		[]float64{4, 0, -2, 1, 671},    //32
		[]float64{2, 0, 0, -3, 607},    //33
		[]float64{2, 0, 2, -1, 596},    //34
		[]float64{2, -1, 1, -1, 491},   //35
		[]float64{2, 0, -2, 1, -451},   //36
		[]float64{0, 0, 3, -1, 439},    //37
		[]float64{2, 0, 2, 1, 422},     //38
		[]float64{2, 0, -3, -1, 421},   //39
		[]float64{2, 1, -1, 1, -366},   //40
		[]float64{2, 1, 0, 1, -351},    //41
		[]float64{4, 0, 0, 1, 331},     //42
		[]float64{2, -1, 1, 1, 315},    //43
		[]float64{2, -2, 0, -1, 302},   //44
		[]float64{0, 0, 1, 3, -283},    //45
		[]float64{2, 1, 1, -1, -229},   //46
		[]float64{1, 1, 0, -1, 223},    //47
		[]float64{1, 1, 0, 1, 223},     //48
		[]float64{0, 1, -2, -1, -220},  //49
		[]float64{2, 1, -1, -1, -220},  //50
		[]float64{1, 0, 1, 1, -185},    //51
		[]float64{2, -1, -2, -1, 181},  //52
		[]float64{0, 1, 2, 1, -177},    //53
		[]float64{4, 0, -2, -1, 176},   //54
		[]float64{4, -1, -1, -1, 166},  //55
		[]float64{1, 0, 1, -1, -164},   //56
		[]float64{4, 0, 1, -1, 132},    //57
		[]float64{1, 0, -1, -1, -119},  //58
		[]float64{4, -1, 0, -1, 115},   //59
		[]float64{2, -2, 0, 1, 107}}    //60

	//TABLE 22.A
	//MEEUS PAGES 144
	//Periodic terms for the nutation in longitude
	sumNutationLongitudeArray = [][]float64{
		[]float64{0, 0, 0, 0, 1, -171996, -174.2, 92025, 8.9}, //1
		[]float64{-2, 0, 0, 2, 2, -13187, -1.6, 5736, -3.1},   //2
		[]float64{0, 0, 0, 2, 2, -2274, -0.2, 977, -0.5},      //3
		[]float64{0, 0, 0, 0, 2, 2062, 0.2, -895, 0.5},        //4
		[]float64{0, 1, 0, 0, 0, 1426, -3.4, 54, -0.1},        //5
		[]float64{0, 0, 1, 0, 0, 712, 0.1, -7, 0},             //6
		[]float64{-2, 1, 0, 2, 2, -517, 1.2, 224, -0.6},       //7
		[]float64{0, 0, 0, 2, 1, -386, -0.4, 200, 0},          //8
		[]float64{0, 0, 1, 2, 2, -301, 0, 129, -0.1},          //9
		[]float64{-2, -1, 0, 2, 2, 217, -0.5, -95, 0.3},       //10
		[]float64{-2, 0, 1, 0, 0, -158, 0, 0, 0},              //11
		[]float64{-2, 0, 0, 2, 1, 129, 0.1, -70, 0},           //12
		[]float64{0, 0, -1, 2, 2, 123, 0, -53, 0},             //13
		[]float64{2, 0, 0, 0, 0, 63, 0, 0, 0},                 //14
		[]float64{0, 0, 1, 0, 1, 63, 0.1, -33, 0},             //15
		[]float64{2, 0, -1, 2, 2, -59, 0, 26, 0},              //16
		[]float64{0, 0, -1, 0, 1, -58, -0.1, 32, 0},           //17
		[]float64{0, 0, 1, 2, 1, -51, 0, 27, 0},               //18
		[]float64{-2, 0, 2, 0, 0, 48, 0, 0, 0},                //19
		[]float64{0, 0, -2, 2, 1, 46, 0, -24, 0},              //20
		[]float64{2, 0, 0, 2, 2, -38, 0, 16, 0},               //21
		[]float64{0, 0, 2, 2, 2, -31, 0, 13, 0},               //22
		[]float64{0, 0, 2, 0, 0, 29, 0, 0, 0},                 //23
		[]float64{-2, 0, 1, 2, 2, 29, 0, -12, 0},              //24
		[]float64{0, 0, 0, 2, 0, 26, 0, 0, 0},                 //25
		[]float64{-2, 0, 0, 2, 0, -22, 0, 0, 0},               //26
		[]float64{0, 0, -1, 2, 1, 21, 0, -10, 0},              //27
		[]float64{0, 2, 0, 0, 0, 17, -0.1, 0, 0},              //28
		[]float64{2, 0, -1, 0, 1, 16, 0, -8, 0},               //29
		[]float64{-2, 2, 0, 2, 2, -16, 0.1, 7, 0},             //30
		[]float64{0, 1, 0, 0, 1, -15, 0, 9, 0},                //31
		[]float64{-2, 0, 1, 0, 1, -13, 0, 7, 0},               //32
		[]float64{0, -1, 0, 0, 1, -12, 0, 6, 0},               //33
		[]float64{0, 0, 2, -2, 0, 11, 0, 0, 0},                //34
		[]float64{2, 0, -1, 2, 1, -10, 0, 5, 0},               //35
		[]float64{2, 0, 1, 2, 2, -8, 0, 3, 0},                 //36
		[]float64{0, 1, 0, 2, 2, 7, 0, -3, 0},                 //37
		[]float64{-2, 1, 1, 0, 0, -7, 0, 0, 0},                //38
		[]float64{0, -1, 0, 2, 2, -7, 0, 3, 0},                //39
		[]float64{2, 0, 0, 2, 1, -7, 0, 3, 0},                 //40
		[]float64{2, 0, 1, 0, 0, 6, 0, 0, 0},                  //41
		[]float64{-2, 0, 2, 2, 2, 6, 0, -3, 0},                //42
		[]float64{-2, 0, 1, 2, 1, 6, 0, -3, 0},                //43
		[]float64{2, 0, -2, 0, 1, -6, 0, 3, 0},                //44
		[]float64{2, 0, 0, 0, 1, -6, 0, 3, 0},                 //45
		[]float64{0, -1, 1, 0, 0, 5, 0, 0, 0},                 //46
		[]float64{-2, -1, 0, 2, 1, -5, 0, 3, 0},               //47
		[]float64{-2, 0, 0, 0, 1, -5, 0, 3, 0},                //48
		[]float64{0, 0, 2, 2, 1, -5, 0, 3, 0},                 //49
		[]float64{-2, 0, 2, 0, 1, 4, 0, 0, 0},                 //50
		[]float64{-2, 1, 0, 2, 1, 4, 0, 0, 0},                 //51
		[]float64{0, 0, 1, -2, 0, 4, 0, 0, 0},                 //52
		[]float64{-1, 0, 1, 0, 0, -4, 0, 0, 0},                //53
		[]float64{-2, 1, 0, 0, 0, -4, 0, 0, 0},                //54
		[]float64{1, 0, 0, 0, 0, -4, 0, 0, 0},                 //55
		[]float64{0, 0, 1, 2, 0, 3, 0, 0, 0},                  //56
		[]float64{0, 0, -2, 2, 2, -3, 0, 0, 0},                //57
		[]float64{-1, -1, 1, 0, 0, -3, 0, 0, 0},               //58
		[]float64{0, 1, 1, 0, 0, -3, 0, 0, 0},                 //59
		[]float64{0, -1, 1, 2, 2, -3, 0, 0, 0},                //60
		[]float64{2, -1, -1, 2, 2, -3, 0, 0, 0},               //61
		[]float64{0, 0, 3, 2, 2, -3, 0, 0, 0},                 //62
		[]float64{2, -1, 0, 2, 2, -3, 0, 0, 0}}                //63
)

// (22.1) Astronomical Algorithms p 143
// Get Julian Day Century from 2000
func getJ2000Century(j2000 float64) float64 {
	return (j2000 - 2451545) / 36525
}

// (47.1) Astronomical Algorithms p 338
// Moon Mean Longitude
func getMoonMeanLongitude(t float64) float64 {
	return (218.3164477 + t*481267.88123421 - t*t*0.0015786 +
		t*t*t/538841 - t*t*t*t/6519400)
}

// (47.2) Astronomical Algorithms p 338
// Moon Mean Elongation
func getMoonMeanElongation(t float64) float64 {
	return (297.8501921 + t*445267.1114034 - t*t*0.0018819 +
		t*t*t/545868 - t*t*t*t/113065000)
}

// (47.4) Astronomical Algorithms p 338
// Moon Mean Anomaly
func getMoonMeanAnomaly(t float64) float64 {
	return (134.9633964 + t*477198.8675055 + t*t*0.0087414 +
		t*t*t/69699 - t*t*t*t/14712000)
}

// (47.5) Astronomical Algorithms p 338
// Moon Argument of Latitude (mean distance of moon from ascending node)
func getMoonArgumentOfLatitude(t float64) float64 {
	return (93.2720950 + t*483202.0175233 - t*t*0.0036539 -
		t*t*t/3526000 + t*t*t*t/863310000)
}

// Astronomical Algorithms p 338
// Moon Correction Term A1 (due to action of Venus)
func getMoonA1(t float64) float64 {
	return 119.75 + t*131.849
}

// Astronomical Algorithms p 338
// Moon Correction Term A2 (due to action of Jupiter)
func getMoonA2(t float64) float64 {
	return 53.09 + t*479264.29
}

// Astronomical Algorithms p 338
// Moon Correction Term A3 (due to flattening of Earth)
func getMoonA3(t float64) float64 {
	return 313.45 + t*481266.484
}

// (47.6) Astronomical Algorithms p 338
// Earth Eccentricity for Moon Calculations
func getEarthEccentricityForMoon(t float64) float64 {
	return 1 - t*0.002516 - t*t*0.0000074
}

// (47.7) Astronomical Algorithms p 338
// Moon Longitude of Ascending Node
func getMoonLongitudeOfAscendingNode(t float64) float64 {
	return (125.04452 - t*1934.136261 + t*t*0.0020708 +
		t*t*t/450000)
}

func calculateSumsLAndR(lp, d, m, mp, f, t float64) (sumL, sumR float64) {
	a1 := getMoonA1(t)
	a1Rad := degreesToRadians(a1)
	a2 := getMoonA2(t)
	a2Rad := degreesToRadians(a2)
	moonEarthEccentricity := getEarthEccentricityForMoon(t)
	for _, x := range sumLRArray {
		termD := x[0] * d
		termM := x[1] * m
		termMP := x[2] * mp
		termF := x[3] * f
		term := termD + termM + termMP + termF
		termRad := degreesToRadians(term)
		termSin := math.Sin(termRad)
		termCos := math.Cos(termRad)
		if x[1] != 0 {
			if x[4] != 0 {
				termSin *= x[4] * moonEarthEccentricity
			} else {
				termSin *= moonEarthEccentricity
			}
			if x[5] != 0 {
				termCos *= x[5] * moonEarthEccentricity
			} else {
				termCos *= moonEarthEccentricity
			}
		} else {
			if x[4] != 0 {
				termSin *= x[4]
			}
			if x[5] != 0 {
				termCos *= x[5]
			}
		}
		sumL += termSin
		sumR += termCos
	}
	sumL += 3958 * math.Sin(a1Rad)
	sumL += 1962 * math.Sin(degreesToRadians(lp-f))
	sumL += 318 * math.Sin(a2Rad)
	return
}

func calculateSumB(lp, d, m, mp, f, t float64) (sumB float64) {
	a1 := getMoonA1(t)
	a3 := getMoonA1(t)
	a3Rad := degreesToRadians(a3)
	moonEarthEccentricity := getEarthEccentricityForMoon(t)
	for _, x := range sumBArray {
		termD := x[0] * d
		termM := x[1] * m
		termMP := x[2] * mp
		termF := x[3] * f
		term := termD + termM + termMP + termF
		term = math.Sin(degreesToRadians(term))
		if x[1] != 0 {
			if x[4] != 0 {
				term *= x[4] * moonEarthEccentricity
			} else {
				term *= moonEarthEccentricity
			}
		} else {
			if x[4] != 0 {
				term *= x[4]
			}
		}
		sumB += term
	}
	sumB += 2235 * math.Sin(degreesToRadians(lp))
	sumB += 382 * math.Sin(a3Rad)
	sumB += 175 * math.Sin(degreesToRadians(a1-f))
	sumB += 175 * math.Sin(degreesToRadians(a1+f))
	sumB += 127 * math.Sin(degreesToRadians(lp-mp))
	sumB += 115 * math.Sin(degreesToRadians(lp+mp))
	return
}

func calculateSumsDeltaYAndE(lp, d, m, mp, f, o, t float64) (sumY, sumE float64) {
	for _, x := range sumNutationLongitudeArray {
		termD := x[0] * d
		termM := x[1] * m
		termMP := x[2] * mp
		termF := x[3] * f
		termOmega := x[4] * o
		term := termD + termM + termMP + termF + termOmega
		termRad := degreesToRadians(term)
		termSin := math.Sin(termRad)
		termCos := math.Cos(termRad)
		if x[6] != 0 {
			termSin *= (x[5] + x[6]*t)
		} else {
			termSin *= x[5]
		}
		if x[7] != 0 {
			if x[8] != 0 {
				termCos *= (x[7] + x[8]*t)
			} else {
				termCos *= x[7]
			}
		}
		sumY += termSin
		sumE += termCos
	}
	sumY *= 0.0001 / 3600
	sumE *= 0.0001 / 3600
	return
}

func degreesToArctime(d float64) (degree, arcMinute, arcSecond float64) {
	modD := modDegrees(d)
	degree, arcTime := math.Modf(modD)
	arcTime *= 3600
	arcMinute, _ = math.Modf(arcTime / 60)
	arcSecond = math.Mod(arcTime, 60)
	return
}

// (48.4) Astronomical Algorithms p 346
func getMoonIlluminationFraction(d, m, mp float64) float64 {
	dDeg := radiansToDegrees(d)
	a := (180 - dDeg - 6.289*math.Sin(mp) + 2.1*math.Sin(m) -
		1.274*math.Sin(2*d-mp) - 0.658*math.Sin(2*d) -
		0.214*math.Sin(2*mp) - 0.11*math.Sin(d))
	return a
}

// (48.1) Astronomical Algorithms p 345
func getIllumination(i float64) float64 {
	return (1 + math.Cos(degreesToRadians(i))) / 2
}
func getMoonLocationIllumination(t time.Time) (eciCoords satellite.Vector3, i float64) {
	year, month, day := t.UTC().Date()
	hour, minute, second := t.UTC().Clock()
	jday := satellite.JDay(year, int(month), day, hour, minute, second)
	jC := getJ2000Century(jday)

	lp := getMoonMeanLongitude(jC)
	d := getMoonMeanElongation(jC)
	dRad := degreesToRadians(d)
	m := getSunMeanAnomaly(jC)
	mRad := degreesToRadians(m)
	mp := getMoonMeanAnomaly(jC)
	mpRad := degreesToRadians(mp)
	f := getMoonArgumentOfLatitude(jC)
	o := getMoonLongitudeOfAscendingNode(jC)
	sumL, sumR := calculateSumsLAndR(lp, d, m, mp, f, jC)
	l := lp + sumL/1000000
	delta := 385000.56 + sumR/1000
	sumY, sumE := calculateSumsDeltaYAndE(lp, d, m, mp, f, o, jC)
	epsilon := 23.43929 - jC*0.01300417 - jC*jC*0.0000001638889*jC*jC*jC*0.0000005036111 + sumE
	epsilonRad := degreesToRadians(epsilon)
	apparentLongitude := l + sumY
	apparentLongitudeRad := degreesToRadians(apparentLongitude)
	eciCoords.X = math.Cos(apparentLongitudeRad) * delta
	eciCoords.Y = math.Cos(epsilonRad) * math.Sin(apparentLongitudeRad) * delta
	eciCoords.Z = math.Sin(epsilonRad) * math.Sin(apparentLongitudeRad) * delta
	i = getMoonIlluminationFraction(dRad, mRad, mpRad)
	return
}
*/
