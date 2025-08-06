import { View, FlatList, Dimensions } from 'react-native';
import React, { useEffect, useState } from 'react';
import LykedCard from '../components/LykedCard';

const CARD_MARGIN = 8;
const NUM_COLUMNS = 2;
const SCREEN_WIDTH = Dimensions.get('window').width;
const CARD_WIDTH = (SCREEN_WIDTH - (NUM_COLUMNS + 1) * CARD_MARGIN) / NUM_COLUMNS;

const API_KEY = 'ad738c4cf4e7548994a8e95ef192859c'; // Replace with real API key
const DEFAULT_IMAGE = 'https://via.placeholder.com/150';

const HomeScreen = () => {
  const [uploads, setUploads] = useState<Upload[]>([]);

  type Upload = {
    id: string;
    title: string;
    description: string;
    video_link: string;
    user_id: string;
    previewImage?: string;
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(
          'http://localhost:8080/uploads/all?user_id=7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df'
        );
        const data = await response.json();

        const withPreviewImages = await Promise.all(
          (data.uploads || []).map(async (upload: Upload) => {
            try {
              const res = await fetch(
                `https://api.linkpreview.net/?key=${API_KEY}&q=${encodeURIComponent(upload.video_link)}`
              );
              const meta = await res.json();
              return {
                ...upload,
                previewImage: meta?.image || DEFAULT_IMAGE,
              };
            } catch {
              return {
                ...upload,
                previewImage: DEFAULT_IMAGE,
              };
            }
          })
        );

        setUploads(withPreviewImages);
      } catch (error) {
        console.error('Failed to fetch uploads:', error);
      }
    };

    fetchData();
  }, []);

  const renderItem = ({ item }: { item: Upload }) => (
    <LykedCard
      title={item.title}
      description={item.description}
      previewImage={item.previewImage || DEFAULT_IMAGE}
      width={CARD_WIDTH}
    />
  );

  return (
    <View className="flex-1 bg-gray-100 p-2">
      <FlatList
        data={uploads}
        renderItem={renderItem}
        keyExtractor={(item) => item.id}
        numColumns={NUM_COLUMNS}
        columnWrapperStyle={{ justifyContent: 'space-between' }}
        contentContainerStyle={{ paddingBottom: 20 }}
      />
    </View>
  );
};

export default HomeScreen;
