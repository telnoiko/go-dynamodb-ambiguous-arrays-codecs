export function testyUnmarshaledValues(response, expectedValues) {
    return response.hasOwnProperty("favorite_food") && // Cannot find 'favorite_food' field in response
        response.propertyIsEnumerable("favorite_food") && //'favorite_food' field is not an array
        expectedValues.every(v => response.favorite_food.includes(v))  // 'favorite_food' field value is broken
}