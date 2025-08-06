import { View, Text } from 'react-native';
import React from 'react';

const Navbar = () => {
  return (
    <View className="h-15 flex-row items-center justify-between border-b border-gray-300 bg-white px-4">
      <Text className="text-lg font-bold">Logo</Text>
      <Text className="text-base">Menu</Text>
    </View>
  );
};

export default Navbar;
