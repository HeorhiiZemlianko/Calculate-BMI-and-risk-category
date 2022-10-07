<h1 align="center">Calculate BMI and risk category
<p target="_blank">IoT <img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></p></h1>
<h3 align="center">This program implements a simple attempt to calculate BMI and critical categories based on the results obtained.
</h3>
<img src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103" >

## The goal of the work
The main goal of this project is to develop a program that calculates BMI and shows the critical category according to the results.

## Task statement
Implement a program that determines whether a person's body weight is normal. A person's weight in kilograms should normally be equal to his height in centimeters minus 100 +/-10%. If the height is more than 10%, then it is necessary to display "needs improvement". If the weight is more than 10%, then you need to display "need to lose weight". If normal, then "normal".

# Ideal weight formula: the most popular calculation methods
- ## Brock's ideal weight formula
One of the most reasonable ways to calculate the ideal weight is the Brock formula. It takes into account the ratio of weight, height and age of a person. With age, the weight of a woman and a man should gradually increase - this is a normal physiological process. And kilograms, which some may consider "superfluous", in fact, may not be. Brock's formula for those who have not reached the age of forty is "height (in cm) - 110", after forty years - "height (in cm) - 100". At the same time, fragile thin people - **asthenics** - must subtract 10% from the result, and owners of a massive physique - **hypersthenics** - must add 10% to it.
- Brock's formula for calculating ideal weight based on age:
``` 
Up to 40 years (kg) = height (cm) - 110
After 40 years (kg) = height (cm) - 100
``` 
After that, asthenics subtract 10%, and hypersthenics add 10%.

- ## Calculation of the ideal weight using the Lorenz formula
Neither the Lorenz formula nor the ideal weight tables take into account body type, which introduces a significant error. Depending on your body type, your ideal weight may vary.
- Lorentz formula (for the calculation you only need to know your height):
```
Ideal weight = (height (cm) - 100) - (height (cm) - 150) / 2
```

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
