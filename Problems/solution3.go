func lengthOfLongestSubstring(s string) int {
  originalString:=[]rune(s)
  var exists bool
  m := make(map[rune]int) //map to store character indices
  longest:=0
  follower:=0
  leader:=0

    for leader<len(originalString){
      _,exists=m[originalString[leader]]
      if (exists){
          delete(m,originalString[follower])
          follower=follower+1
      }else{
          m[originalString[leader]]=leader
          leader=leader+1
          if(len(m)>longest){longest=len(m)}
      }
  }
  return longest
}
