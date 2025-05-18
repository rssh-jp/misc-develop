import { NextResponse, NextRequest } from "next/server";
import { ELATICSEARCH_HOSTS } from "@/internal/const/common";

export async function POST(req: NextRequest) {
  const body = await req.json();
  console.log("++++++++++", req, body);
  // ElasticsearchのエンドポイントURL
  const elasticsearchUrl = ELATICSEARCH_HOSTS + "/wiki/_search";

  // 固定の検索クエリボディ (ユーザー指定のもの)
  const queryBody = {
    query: {
      bool: {
        must: [
          {
            multi_match: {
              query: "あああ", // 将来的には動的に変更可能
              fields: ["title", "text"],
            },
          },
        ],
      },
    },
    size: 1,
    sort: {
      _score: {
        order: "desc",
      },
    },
  };

  try {
    console.log(
      "Sending query to Elasticsearch:",
      JSON.stringify(queryBody, null, 2),
    );

    const esResponse = await fetch(elasticsearchUrl, {
      method: "POST", // GETでもbodyを送れるが、POSTが一般的
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(queryBody),
    });

    // Elasticsearchからのレスポンスをチェック
    if (!esResponse.ok) {
      // エラー詳細を取得試行
      let errorBody = null;
      try {
        errorBody = await esResponse.json();
        console.error("Elasticsearch error response:", errorBody);
      } catch (e) {
        console.error(
          "Could not parse Elasticsearch error response:",
          await esResponse.text(),
        );
      }
      throw new Error(
        `Elasticsearch query failed with status ${esResponse.status}`,
      );
    }

    // レスポンスボディ (JSON) を取得
    const data = await esResponse.json();

    return NextResponse.json({ message: data }, { status: 200 });
  } catch (error) {
    console.error("Error contacting Elasticsearch:", error);
    return NextResponse.json({ message: error.message }, { status: 500 });
  }
}
