export const solution = (arr1: number[], arr2: number[]) => {
  let similarity = 0;

  arr1.forEach((value, index) => {
    similarity += arr2.filter((v) => v === value).length * value;
  });

  return similarity;
};