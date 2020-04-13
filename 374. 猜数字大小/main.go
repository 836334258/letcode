package main

func guessNumber(n int) int {
	i:=1
	j:=n
	for i<=j{
		t:=i+(j-i)/2
		if guess(t)==0 {
			return t
		}
		if guess(t)==1 {
			i=t+1
		}else{
			j=t-1
		}
	}
	return -1
}

func main()  {

}
