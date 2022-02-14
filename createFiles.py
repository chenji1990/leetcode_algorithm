import os

def createFiles(folderPath: str):
    for i in range(5, 53):
        path = folderPath + "{:0>4d}".format(i)
        if not os.path.exists(path):
            os.mkdir(path)
        if not os.path.exists(f"{path}/main.go"):
            with open(f"{path}/main.go", "w", encoding="utf-8") as file:
                file.write("package main\n\n")        
        if not os.path.exists(f"{path}/main.py"):
            with open(f"{path}/main.py", "w", encoding="utf-8") as file:
                file.write("\n\n")

createFiles("./Offer2/")