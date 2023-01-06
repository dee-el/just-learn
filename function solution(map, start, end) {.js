function solution(map, start, end) {
    if (map[start[0]] === undefined || map[start[1]] === undefined) {
      return 0;
    }
  
    let row_length = map.length;
    let col_length = map[0].length;
  
    let visited = new Set();
  
    let queue = [];
  
    queue.push({coord: start, steps: 1});
    visited.add(map[start[0]][start[1]].toString());
  
    // BOTTOM UP LEFT RIGHT 
    let d_row = [-1, 1, 0, 0];
    let d_col = [0, 0, -1, 1];
  
  
    while (queue.length > 0) {
      let curr = queue.shift();
  
      let row = curr.coord[0];
      let col = curr.coord[1];
  
      if (row === end[0] && col === end[1]) {
     
        return curr.steps
      }
  
      // BOTTOM UP LEFT RIGHT 
      for (let mv = 0; mv < 4; mv++) {
        let t_row = row + d_row[mv];
        let t_col = col + d_col[mv];
  
        if (
          !visited.has([t_row, t_col].toString()) && 
          t_row >= 0 &&
          t_col >= 0 &&
          t_row < row_length &&
          t_col < col_length &&
          map[t_row][t_col] !== map[row][col]
        ) {
          visited.add([t_row, t_col].toString())
  
          queue.push({coord: [t_row, t_col], steps: curr.steps+1});
        }
      }
    }
  
    return 0;
  }
  exports.solution = solution;
  
  