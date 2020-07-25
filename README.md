## Dynamic Programming - Maximum Height at Each Point
* Finding the maximum height of the rectangle by iterating upwards until a 0 is reached
* Finding the maximum width of the rectangle by iterating outwards left and right until a height that doesn't accommodate the maximum height of the rectangle

keep 3 lists. Height, Left and Right. Loop through the rows. Fill in each
Keep track of maxArea

loop over rows:

* initialize current_left as 0
* initialize current_right as height-1

Loop over columns to calculate height, left and right

If there is a "1":
  * Height: height[j]++
  * Left: max(left[j],curr_left)
  * Right: min(right[j],curr_right)

if there is a "0":
  * Height: 0
  * Left: 
    * Left[column] = 0
    * current_left = (column + 1)
  * Right:
    * Right[column] = height - 1
    * current_right = (column - 1)


Area: 
  * calculate current:  (right[col] - left[col] + 1) * height[col]
  * maxArea = max(maxArea, current)

Time complexity O(N*M)
Space: O(M)

