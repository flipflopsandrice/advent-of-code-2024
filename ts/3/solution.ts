export const solution = (input: string) => {
  const matches = [...input.matchAll(/mul\(([0-9+]{0,}),([0-9+]{0,})\)/gm)];

  let total = 0;
  (matches ?? []).forEach((match) => {
    const [, a, b] = match;
    total += parseInt(a) * parseInt(b);
  });

  return total;
};

export const part2 = (input: string) => {
  const matches = [...input.matchAll(/(do|don\'t|mul)\((([0-9+]{0,}),([0-9+]{0,}))?\)/gm)];

  let total = 0;
  let active = true;
  (matches ?? []).forEach((match) => {
    const [, func, , a, b] = match;
    if (func === 'mul' && active) {
      total += parseInt(a) * parseInt(b);
    } else {
      active = func === 'do'
    }
  });

  return total;
};