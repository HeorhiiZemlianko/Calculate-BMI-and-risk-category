<h1 align="center">Calculate BMI and risk category
<p target="_blank">IoT <img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></p></h1>
<h3 align="center">This program implements a simple attempt to calculate BMI and critical categories based on the results obtained.
</h3>
<img src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103" >

## The goal of the work
The main goal of this project is to develop a program that calculates BMI and shows the critical category according to the results.

## Task statement
Implement a program that determines whether a person's body weight is normal. A person's weight in kilograms should normally be equal to his height in centimeters minus 100 +/-10%. If the height is more than 10%, then it is necessary to display "needs improvement". If the weight is more than 10%, then you need to display "need to lose weight". If normal, then "normal".

## Notes
This program is implemented in the **Go** language, the following libraries were used: **`fmt`** and **`math`**

## Notation of variables
This table shows the short designation of the variable in the code and its description

| Name       | Type   | Description                      |
| ---------- | ------ | -------------------------------- |
| `weight` | float64 | the weight |
| `height` | float64 | the growth  |
| `bmi` | float64 | Body mass index (Quetelet index) |

## Functions in the project code

- Calculate BMI and convert cm to meters
``` Go
heightInMeter := height / 100
bmi = weight / math.Pow(heightInMeter, 2)
```

- Category validation conditions
``` Go
if weight < ((height - 100) - (heightInMeter * 10) + 10) {
		fmt.Println("Нужно поправиться.")
} else if weight > ((height - 100) + (heightInMeter * 10)) {
		fmt.Println("Нужно похудеть.")
} else {
		fmt.Println("Нормальный вес.")
}
```

- Conditions for checking weight by BMI
``` Go
normalWeight := (height - 100) - (height-150)/2
delta := weight - normalWeight
fmt.Printf("Идеальный вес для вашего роста: %0.2v кг.\n", normalWeight)

if (delta > 0) && (bmi > 30) {
		fmt.Printf("Вам нужно похудеть до идеального веса на %0.2v кг.\n", math.Abs(delta))
} else if (delta > 0) && (bmi > 25) {
		fmt.Printf("Вам нужно похудеть до идеального веса на %0.2v кг.\n", math.Abs(delta))
} else if (delta < 0) && (bmi < float64(18.5)) {
		fmt.Printf("Вам нужно поправится до идеального веса на %0.2v кг.\n", math.Abs(delta))
}
```
## Result
Screenshot that displays the result of the entered data
<p align="center">
<img  src="https://github.com/HeorhiiZemlianko/Calculate-BMI-and-risk-category/blob/main/task2golang/Снимок.PNG"  width="350" alt="Calculate-BMI-and-risk-category"/>
</p>
