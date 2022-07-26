var exist = function (board, word) {
  const height = board.length
  const width = board[0].length
  let cacheChar;

  function walk(x, y, index) {
    if (index === word.length - 1 && board[y][x] === word[index]) {
      return true;
    }

    cacheChar = board[y][x]

    if (board[y][x] === word[index]) {

      board[y][x] = ''
      if (x + 1 < width && walk(x + 1, y, index + 1)) {
        return true;
      }
      if (x - 1 >= 0 && walk(x - 1, y, index + 1)) {
        return true;
      }

      if (y - 1 >= 0 && walk(x, y - 1, index + 1)) {
        return true;
      }
      if (y + 1 < height && walk(x, y + 1, index + 1)) {
        return true;
      }
    }

    board[y][x] = cacheChar
    return false
  }

  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      if (walk(x, y, 0)) {
        return true
      }
    }
  }

  return false
};