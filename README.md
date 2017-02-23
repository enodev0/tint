# tint

tint performs numerical integration using trapezium rule. This is the single threaded version.

## Usage

Suppose you want to integrate the function f(x) = sin(1/(1-(x**2))) from 0 to PI, you would do:
```bash
$ tintp -ll=0 -ul=3.14 -fn="sin(1/(1-(x**2)))" -n=1200
```
where n = #trapezoids you want to use to calculate the integral.

The greater the value of *n*, the more accurate but resource intensive the calculation will be.

## Caveat

* This is faster for comparatively simpler integrations on <=4 core machines with limited memory, compared to tintp.
  
* Using very high values of *n* may quickly deplete your system resources. Your machine may freeze. Exercise caution.
