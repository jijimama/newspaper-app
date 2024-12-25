import React from 'react';

type Article = {
  id: number;
  year: number;
  month: number;
  day: number;
  content: string;
  newspaper: string;
  column_name: string;
};

async function fetchArticles(): Promise<Article[]> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  console.log('API URL:', apiUrl); // 環境変数を確認するために追加
  if (!apiUrl) {
    throw new Error('NEXT_PUBLIC_API_URL is not defined');
  }
  try {
    const res = await fetch(`${apiUrl}/articles`, { cache: 'no-store' });
    if (!res.ok) {
      const errorText = await res.text();
      throw new Error(`Failed to fetch articles: ${res.status} ${res.statusText} - ${errorText}`);
    }
    return res.json();
  } catch (error) {
    console.error('Error fetching articles:', error);
    throw error;
  }
}

const ArticlesPage = async () => {
  const articles = await fetchArticles();

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-center text-2xl font-bold mb-4">コラム一覧</h1>
      <ul className="space-y-4">
        {articles.map((article) => (
          <li key={article.id} className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-lg font-semibold text-gray-300 mb-2">{article.content}</h2>
            <p className="text-gray-400">{article.year}/{article.month}/{article.day} {article.newspaper} / {article.column_name}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ArticlesPage;
