export function testyUnmarshaledValues(response, expectedValues) {
    return response.hasOwnProperty("choice") && // Cannot find 'choice' field in response
        response.propertyIsEnumerable("choice") && //'choice' field is not an array
        expectedValues.every(v => response.choice.includes(v))  // 'choice' field value is broken
}