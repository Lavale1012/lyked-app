import { View, Text, Image, ActivityIndicator } from 'react-native';
// import React, { useState, useEffect } from 'react';

// const DEFAULT_IMAGE = 'https://via.placeholder.com/150';
// Replace with your real key

const LykedCard = ({
  title,
  description,
  previewImage,
  width,
}: {
  title: string;
  description: string;
  previewImage: string;
  width: number;
}) => {
  return (
    <View
      style={{ width, marginBottom: 12 }}
      className="overflow-hidden rounded-lg bg-white shadow-sm">
      {previewImage ? (
        <Image
          source={{ uri: previewImage }}
          style={{ width: '100%', height: width * 0.6 }}
          resizeMode="cover"
        />
      ) : (
        <View
          style={{ width: '100%', height: width * 0.6 }}
          className="items-center justify-center">
          <ActivityIndicator size="small" color="#888" />
        </View>
      )}

      <View className="p-2">
        <Text className="text-sm font-semibold">{title}</Text>
        <Text className="text-xs text-gray-600">{description}</Text>
      </View>
    </View>
  );
};

export default LykedCard;
