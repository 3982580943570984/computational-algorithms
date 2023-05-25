#include <cmath>
#include <iostream>

using namespace std;

const double a = 0.0;
const double b = 1.0;
const int n = 10;
const double e = 1e-8;

double f(double x) { return exp(-x) * cos(pow(x, 2)); }

double IntegrateLeftRectangles(double a, double b, int n) {
  double h = (b - a) / double(n);
  double integral = 0.0;

  for (int i = 0; i < n; i++) {
    double x = a + double(i) * h;
    integral += f(x) * h;
  }

  return integral;
}

double IntegrateRightRectangles(double a, double b, int n) {
  double h = (b - a) / double(n);
  double integral = 0.0;

  for (int i = 0; i < n; i++) {
    double x = a + double(i + 1) * h;
    integral += f(x) * h;
  }

  return integral;
}

double IntegrateMidRectangles(double a, double b, int n) {
  double h = (b - a) / double(n);
  double integral = 0.0;

  for (int i = 0; i < n; i++) {
    double x = a + double(i) * h + h / 2.0;
    integral += f(x) * h;
  }

  return integral;
}

double IntegrateTrapezoid(double a, double b, int n) {
  double h = (b - a) / double(n);
  double integral = (f(a) + f(b)) / 2.0;

  for (int i = 1; i < n; i++) {
    double x = a + double(i) * h;
    integral += f(x);
  }

  integral *= h;
  return integral;
}

double IntegrateSimpson(double a, double b, int n) {
  double h = (b - a) / double(n);
  double integral = f(a) + f(b);

  for (int i = 1; i < n; i++) {
    double x = a + double(i) * h;
    if (i % 2 == 0) {
      integral += 2.0 * f(x);
    } else {
      integral += 4.0 * f(x);
    }
  }

  integral *= h / 3.0;
  return integral;
}

// https://ru.wikipedia.org/wiki/Правило_Рунге
double RungeRule(double I1, double I2, double O) { return abs(I2 - I1) / O; }

int main() {
  int multiplier;
  double I1, I2;

  // Метод прямоугольников
  multiplier = 2;
  I1 = IntegrateLeftRectangles(a, b, n);
  I2 = IntegrateLeftRectangles(a, b, multiplier * n);
  while (RungeRule(I1, I2, 1.0) > e) {
    multiplier *= 2;
    I1 = I2;
    I2 = IntegrateLeftRectangles(a, b, multiplier * n);
  }
  printf("Значение интеграла: %.30f\n", I2);

  // Метод правых прямоугольников
  multiplier = 2;
  I1 = IntegrateRightRectangles(a, b, n);
  I2 = IntegrateRightRectangles(a, b, multiplier * n);
  while (RungeRule(I1, I2, 1.0) > e) {
    multiplier *= 2;
    I1 = I2;
    I2 = IntegrateRightRectangles(a, b, multiplier * n);
  }
  printf("Значение интеграла: %.30f\n", I2);

  // Метод средних прямоугольников
  multiplier = 2;
  I1 = IntegrateMidRectangles(a, b, n);
  while (RungeRule(I1, I2, 3.0) > e) {
    multiplier *= 2;
    I1 = I2;
    I2 = IntegrateMidRectangles(a, b, multiplier * n);
  }
  printf("Значение интеграла: %.30f\n", I2);

  // Метод трапеций
  multiplier = 2;
  I1 = IntegrateTrapezoid(a, b, n);
  I2 = IntegrateTrapezoid(a, b, multiplier * n);
  while (RungeRule(I1, I2, 3.0) > e) {
    multiplier *= 2;
    I1 = I2;
    I2 = IntegrateTrapezoid(a, b, multiplier * n);
  }
  printf("Значение интеграла: %.30f\n", I2);

  // Метод Симпсона
  multiplier = 2;
  I1 = IntegrateSimpson(a, b, n);
  I2 = IntegrateSimpson(a, b, multiplier * n);
  while (RungeRule(I1, I2, 15.0) > e) {
    multiplier *= 2;
    I1 = I2;
    I2 = IntegrateSimpson(a, b, multiplier * n);
  }
  printf("Значение интеграла: %.30f\n", I2);

  return 0;
}