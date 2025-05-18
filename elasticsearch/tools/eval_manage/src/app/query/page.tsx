"use client";
import { useState, useEffect, useRef } from "react";
import { Button, Box, Grid, GridItem, Textarea, Text } from "@chakra-ui/react";
import TextEditable from "@/components/TextEditable/TextEditable"; // 追加: TextEditableコンポーネントをインポート

export default function QueryPage() {
  const [query, setQuery] = useState("");
  const [result, setResult] = useState("");
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const [textAreaHeight, setTextAreaHeight] = useState("auto");
  const [isLoading, setIsLoading] = useState(false);

  // 入力値変更時に高さを調整
  useEffect(() => {
    if (textAreaRef.current) {
      setTextAreaHeight(`${textAreaRef.current.scrollHeight}px`);
    }
  }, [query]);

  const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setQuery(event.target.value);
  };

  const handleSearch = async () => {
    setIsLoading(true);
    try {
      // Next.js の API ルートを呼び出す
      const response = await fetch("/api/query", {
        method: "POST", // APIルートで設定したメソッドに合わせる
        headers: {
          "Content-Type": "application/json", // なくても良い場合が多いが念のため
        },
        body: JSON.stringify({ aaa: "nice", bbb: 20 }),
        // ボディはAPIルート側で固定されているのでここでは送らない
        // もし動的にしたい場合は body: JSON.stringify({ searchTerm: 'ぐるなび' }) などで送る
      });
      console.log("+++++", response);

      if (!response.ok) {
        // APIルートからのエラーレスポンスを取得
        const errorData = await response.json();
        throw new Error(
          errorData.message ||
            `API request failed with status ${response.status}`,
        );
      }

      const data = await response.json();
      setResult(JSON.stringify(data)); // 結果をstateに保存
    } catch (err) {
      console.error("Failed to fetch from API route:", err);
      setResult(err.message);
    } finally {
      setIsLoading(false); // ローディング完了
    }
  };

  return (
    <Grid templateColumns="repeat(11, 1fr)" gap={4} height="100%" padding="4">
      {/* 左側：クエリ入力エリア */}
      <GridItem colSpan={5}>
        <Box
          padding="0"
          height="100%"
          overflow="hidden"
        >
          {/* クエリ入力欄 */}
          <TextEditable
            query={query}
            onChange={handleInputChange}
            title="クエリ"
          />
        </Box>
      </GridItem>

      {/* 仕切り */}
      <GridItem colStart={6} colSpan={1}>
        <Button onClick={handleSearch} loading={isLoading}>
          検索
        </Button>
        :
      </GridItem>

      {/* 右側：結果表示 */}
      <GridItem colStart={7} colSpan={5}>
        <Box
          border="1px solid"
          borderColor="gray.300"
          padding="4"
          height="100%"
          overflow="auto"
        >
          <Text fontSize="md" fontFamily="monospace">
            {result}
          </Text>
        </Box>
      </GridItem>
    </Grid>
  );
}
