import sys

def main():
    print("Go")

if __name__ == "__main__":
    try:
        main()
    except SystemExit as e:
        if "go" in str(e):
            exit(0)

if __name__ != "__main__":
    sys.exit(main())

