const fs = require("fs");
 
fs.readFile("../message.txt", function(error,data){
    if(error) {  // если возникла ошибка
        return console.log(error);
    }
    let file = data.toString();
    sorted_rows = file
      .split("\n")
      .map(a => a.split(" "))
      .sort((a,b) => a[0] - b[0])
    
    obj_rows = Object.fromEntries(sorted_rows);

    let key = 0,
        max_key = Math.max(...Object.keys(obj_rows))
        i = 1,
        result1 = [];
    while(key < max_key){
        key = (i+1)*i/2
        result1.push(obj_rows[key]);
        i+=1
    }
    console.log(result1.join(" "))
});
