Process:

I initially sparred with chatGPT over a good file structure for such an application. 
I started on the assumption of having to implement a solution using a 2d slice to represent a chessboard where each value has a color assigned to it; this proved to be a naive implementation since
I then remembered that there was a very easy computation to determine whether two values reside on the same diagonal in a matrix. 
After checking what it was (checking if their absolute values are equal), it followed that if focusing solely on two pieces, and only on checking whether a bishop can attack another piece, then only that bit of logic was needed for the application.
(I have left in a board file and minimal logic to include potential avenues of extensibility) 
I have placed more importance on testing of the input validation logic over the 'attack' function, since I wanted to make sure of catching problematic input before calling it. 
(this way, there is not much to test over the logic of the attack function, it being the property of a matrix. In other words: if I can prove that the input received will be correct, I won't have to test each single diagonal and combination on the 'attack' function).
I haven't tested the main input function, mostly because I consider that outside of the scope of the assignment.
