export const distanceSum = (arr1, arr2) => {
    const s1 = arr1.sort();
    const s2 = arr2.sort();

    let distance = 0;
    for (let i = 0; i < s1.length; i++) {
        const v1 = s1[i];
        const v2 = s2[i];

        if (v1 < v2) {
            distance += v2 - v1;
        } else {
            distance += v1 - v2;
        }
    }

    return distance
}