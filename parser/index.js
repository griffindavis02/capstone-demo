fs = require('fs')

let Output = {
    Data: [{
            Rate: 0.5,
            FlipData: [],
        },
        {
            Rate: 0.25,
            FlipData: [],
        },
        {
            Rate: 0.1,
            FlipData: [],
        },
        {
            Rate: 0.01,
            FlipData: [],
        },
    ],
}

fs.readFile('../exampleOutput.json', 'utf8', (err, data) => {
    if (err) {
        console.error(err)
        return
    }

    jsonData = JSON.parse(data)

    for (iteration of jsonData.Data) {
        for (let i = 0; i < Output.Data.length; i++) {
            if (Output.Data[i].Rate == iteration.Rate) {
                Output.Data[i].FlipData.push({
                    IterationNum: iteration.IterationNum,
                    ErrorData: iteration.ErrorData,
                })
            }
        }
    }

    fs.writeFile(
        '../exampleOutputFinal.json',
        JSON.stringify(Output),
        (err) => {
            if (err) {
                console.error(err)
                return
            }
        }
    )
})