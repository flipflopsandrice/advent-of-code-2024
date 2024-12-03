export const solution = (input: string) => {
  const matches = [...input.matchAll(/mul\(([0-9+]{0,}),([0-9+]{0,})\)/gm)];

  let total = 0;
  (matches ?? []).forEach((match) => {
    const [,a, b] = match;
    total += parseInt(a) * parseInt(b);
  });

    return total;
};