git remote - local repo ni asosiy repo bilan bog'lash uchun ishlatiladigan masofadan boshqarish buyrug'i.
git remote -v -->  biron bir masofaviy omborning sozlanganligini tekshirish uchun ishlatiladi
origin - bu faqat havola nomi yoki git omborining URL manzili uchun o'zgaruvchan nom

git remote add istalgannom https://github.com/BSardorbek/lessons1.git
git remote -v
natija:
    istalgannom     https://github.com/BSardorbek/lessons1.git (fetch)
    istalgannom     https://github.com/BSardorbek/lessons1.git (push)
endi biz url manzil o'rniga berilgan nomni ishlatishimiz mumkun

git fetch istalgannom --> bu buiyruq orqali repo da joylashgan ammo bizning local repoda yo'q bo'lgan malumotlarni yuklab oladi

git remote rm istalgannom --> ochirib tashlash
git remote rename istalgannom yaxshiroq nom --> nomni o'zgartirish uchun ishlatiladi
