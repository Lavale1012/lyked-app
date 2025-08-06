import { View } from 'react-native';
import React, { useState } from 'react';
import { TextInput } from 'react-native-gesture-handler';

const SearchBar = () => {
  const [text, setText] = useState('');
  return (
    <View className="h-12 w-full flex-row items-center rounded-xl bg-gray-200 px-4 shadow-sm dark:bg-gray-700">
      <TextInput
        value={text}
        onChangeText={(newText) => setText(newText)}
        placeholder="Search..."
      />
    </View>
  );
};

export default SearchBar;
