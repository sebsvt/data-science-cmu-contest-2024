# Average Score

def main():
	data = []

	while True:
		x = input("data ->")
		if x == "q":
			break
		num = float(x)
		if num >= 0 and num <= 100:
			data.append(float(x))
	avg = sum(data) / len(data)
	if avg >= 50:
		print(avg,"Satisfactory")
	else:
		print(avg,"Unsatisfactory")

if __name__ == "__main__":
	main()
