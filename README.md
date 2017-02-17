# tint

tint performs numerical integration using trapezium rule.

## Usage

Suppose you want to integrate the function f(x) = sin(1/(1-(x**2))) from 0 to PI, you would do:
```bash
$ tint -ll=0 -ul=3.14 -fn="sin(x)" -n=12000
```
In the above case, we specified lower limit(ll) as 0, upper limit(ul) as 3.14, and the function to be 
integrated(fn) as *sin(1/(1-(x**2)))*. Switch n specifies the number of quadrilaterals used to divide up and
approximate the area under the curve. Default value for n is 10.

For something like f(x) = 7, n = 1 will suffice. 