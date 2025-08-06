import { StatusBar } from 'expo-status-bar';
import { Text, View } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import '../../global.css';

import HomeScreen from './screens/HomeScreen';
import FolderScreen from './screens/FolderScreen';
import ProfileScreen from './screens/ProfileScreen';
import CreateLykedScreen from './screens/CreateLykedScreen';
import { GestureHandlerRootView } from 'react-native-gesture-handler';

const Tab = createBottomTabNavigator();

export default function App() {
  return (
    <GestureHandlerRootView>
      <NavigationContainer>
        <View className="flex-1">
          <Tab.Navigator
            screenOptions={{
              headerShown: true,
              // headerRight: () => (
              //   <View className="mr-4">
              //     <Text className="text-lg font-bold">Icon</Text>
              //   </View>
              // ),
              tabBarStyle: {
                flexDirection: 'row',
                paddingTop: 10,

                borderRadius: 50,
                position: 'absolute',
                // overflow: 'hidden',
                marginHorizontal: 10,
                marginBottom: 20,
                // backgroundColor: 'rgba(17, 24, 39, 0.8)',
                elevation: 5, // shadow on Android
                shadowColor: '#000',
                shadowOffset: { width: 3, height: 3 },
                shadowOpacity: 0.2,
                shadowRadius: 10,
              },
            }}>
            <Tab.Screen
              name="Home"
              component={HomeScreen}
              options={{
                title: 'Home',
                headerRight: () => (
                  <View className="mr-4">
                    <Text className="text-lg font-bold">Icon</Text>
                  </View>
                ),
              }}
            />
            <Tab.Screen name="Create" component={CreateLykedScreen} />
            <Tab.Screen name="Folders" component={FolderScreen} />
            <Tab.Screen name="Profile" component={ProfileScreen} />
          </Tab.Navigator>
        </View>
      </NavigationContainer>
      <StatusBar style="auto" />
    </GestureHandlerRootView>
  );
}
