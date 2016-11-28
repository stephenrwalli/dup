# A simple test script to ensure dup works. Includes an empty file.

echo "We are expecting the following numbers:"
cat test_expected.txt
echo " "
echo "Go is not required to maintain maps in the same order,"
echo "so different runs can produce results ordered differently."
echo "Test run output:"
dup testdata/* 


