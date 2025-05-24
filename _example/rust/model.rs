

use serde::{Deserialize, Serialize};
use std::iter::Map;





/**
 * 
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioCouponUnlockForm {
    /**
     * 
     * 
     */
    pub couponId: Option<u64>,
    /**
     * 
     * 
     */
    pub packageId: Option<u64>,
    
}




/**
 * 会员产品
 * 
 */
#[derive(Debug, Clone)]
pub struct VipProductEntity {
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 备注
     * 
     */
    pub description: Option<String>,
    /**
     * 时长
     * 
     */
    pub duration: Option<String>,
    /**
     * 启用标识
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 封面图
     * 
     */
    pub image: Option<String>,
    /**
     * 划线价
     * 
     */
    pub listPrice: Option<usize>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 编码
     * 
     */
    pub no: Option<String>,
    /**
     * 售价
     * 
     */
    pub price: Option<usize>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 标题
     * 
     */
    pub title: Option<String>,
    /**
     * 可用时长（秒）
     * 
     */
    pub token: Option<u64>,
    /**
     * 类型
     * 
     */
    pub type: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct AdNextDayLaunchResult {
    /**
     * 
     * 
     */
    pub icon: Option<String>,
    /**
     * 标识
     * 
     */
    pub id: Option<u64>,
    /**
     * 
     * 
     */
    pub jumpUrl: Option<String>,
    /**
     * 
     * 
     */
    pub launchUrl: Option<String>,
    /**
     * 
     * 
     */
    pub name: Option<String>,
    /**
     * 值
     * 
     */
    pub value: Option<String>,
    
}




/**
 * 视频
 * 
 */
#[derive(Debug, Clone)]
pub struct VideoItem {
    /**
     * 封面图
     * 
     */
    pub image: Option<String>,
    /**
     * 跳转地址
     * 
     */
    pub jumpUrl: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 视频地址
     * 
     */
    pub video: Option<String>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct UserVipRechargeAccount {
    /**
     * 可用语音包数量
     * 
     */
    pub availableAudioCount: Option<String>,
    /**
     * 可用时长(秒)
     * 
     */
    pub availableToken: Option<u64>,
    /**
     * 到期时间
     * 
     */
    pub expireTime: Option<String>,
    /**
     * 
     * 
     */
    pub totalAmount: Option<usize>,
    /**
     * 
     * 
     */
    pub totalCount: Option<u64>,
    /**
     * 状态
     * 
     */
    pub valid: Option<usize>,
    /**
     * 状态文案
     * 
     */
    pub validDescription: Option<String>,
    
}




/**
 * 用户
 * 
 */
#[derive(Debug, Clone)]
pub struct UserItem {
    /**
     * 头像地址
     * 
     */
    pub avatar: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 下载量 / 收益
     * 
     */
    pub weight: Option<usize>,
    
}




/**
 * 用户邀请记录
 * 
 */
#[derive(Debug, Clone)]
pub struct UserInviteLog {
    /**
     * 邀请人头像
     * 
     */
    pub avatar: Option<String>,
    /**
     * 邀请时间
     * 
     */
    pub inviteTime: Option<String>,
    /**
     * 邀请人名称
     * 
     */
    pub name: Option<String>,
    /**
     * 邀请人ID
     * 
     */
    pub userId: Option<u64>,
    
}




/**
 * 用户初始化基础信息
 * 
 */
#[derive(Debug, Clone)]
pub struct UserInitProfile {
    /**
     * 头像
     * 
     */
    pub avatar: Option<String>,
    /**
     * 生日
     * 
     */
    pub birthday: Option<String>,
    /**
     * 性别
     * 
     */
    pub gender: Option<usize>,
    /**
     * 邀请码
     * 
     */
    pub inviteCode: Option<String>,
    /**
     * 昵称
     * 
     */
    pub name: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultBoolean {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<bool>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 用户成长状态
 * 
 */
#[derive(Debug, Clone)]
pub struct UserGrowupState {
    /**
     * 当前经验值
     * 
     */
    pub currentEmpiric: Option<usize>,
    /**
     * 等级
     * 
     */
    pub level: Option<usize>,
    /**
     * 注册天数
     * 
     */
    pub registerDay: Option<usize>,
    /**
     * 累计经验值
     * 
     */
    pub totalEmpiric: Option<usize>,
    
}




/**
 * 用户资金明细
 * 
 */
#[derive(Debug, Clone)]
pub struct UserFinanceEntryDetail {
    /**
     * 交易金额
     * 
     */
    pub transAmount: Option<usize>,
    /**
     * 交易备注
     * 
     */
    pub transRemark: Option<String>,
    /**
     * 交易时间
     * 
     */
    pub transTime: Option<String>,
    /**
     * 交易类型
     * 
     */
    pub transType: Option<String>,
    /**
     * 用户头像
     * 
     */
    pub userAvatar: Option<String>,
    /**
     * 用户ID
     * 
     */
    pub userId: Option<u64>,
    /**
     * 用户昵称
     * 
     */
    pub userName: Option<String>,
    
}




/**
 * 用户账户信息
 * 
 */
#[derive(Debug, Clone)]
pub struct UserFinanceAccount {
    /**
     * 当前申请提现
     * 
     */
    pub applyCash: Option<usize>,
    /**
     * 语音包奖励
     * 
     */
    pub audioReward: Option<usize>,
    /**
     * 余额
     * 
     */
    pub balance: Option<usize>,
    /**
     * 任务奖励
     * 
     */
    pub jobReward: Option<usize>,
    /**
     * 平台奖励
     * 
     */
    pub platformReward: Option<usize>,
    /**
     * 累计提现
     * 
     */
    pub totalCash: Option<usize>,
    /**
     * 累计收益
     * 
     */
    pub totalProfit: Option<usize>,
    /**
     * 会员奖励
     * 
     */
    pub vipReward: Option<usize>,
    
}




/**
 * 行为统计
 * 
 */
#[derive(Debug, Clone)]
pub struct UserBehaviorReference {
    /**
     * 评论
     * 
     */
    pub comment: Option<bool>,
    /**
     * 踩
     * 
     */
    pub dislike: Option<bool>,
    /**
     * 收藏
     * 
     */
    pub favorite: Option<bool>,
    /**
     * 点赞
     * 
     */
    pub like: Option<bool>,
    /**
     * 分享
     * 
     */
    pub share: Option<bool>,
    /**
     * 浏览
     * 
     */
    pub view: Option<bool>,
    
}




/**
 * 升级
 * 
 */
#[derive(Debug, Clone)]
pub struct UpgradeResult {
    /**
     * 是否弹窗
     * 
     */
    pub alter: Option<bool>,
    /**
     * 更新说明
     * 
     */
    pub description: Option<String>,
    /**
     * 下载地址
     * 
     */
    pub download: Option<String>,
    /**
     * 是否强制升级
     * 
     */
    pub focus: Option<bool>,
    /**
     * 当前版本
     * 
     */
    pub version: Option<String>,
    
}




/**
 * 数量统计
 * 
 */
#[derive(Debug, Clone)]
pub struct StatisticReference {
    /**
     * 评论数
     * 
     */
    pub commentCount: Option<u64>,
    /**
     * 不喜欢数
     * 
     */
    pub dislikeCount: Option<u64>,
    /**
     * 收藏数
     * 
     */
    pub favoriteCount: Option<u64>,
    /**
     * 点赞数
     * 
     */
    pub likeCount: Option<u64>,
    /**
     * 分享数
     * 
     */
    pub shareCount: Option<u64>,
    /**
     * 浏览数
     * 
     */
    pub viewCount: Option<u64>,
    
}




/**
 * 奖励提示
 * 
 */
#[derive(Debug, Clone)]
pub struct RewardTip {
    /**
     * 
     * 
     */
    pub content: Option<String>,
    /**
     * 
     * 
     */
    pub progress: Option<usize>,
    /**
     * 
     * 
     */
    pub totalContent: Option<String>,
    
}




/**
 * 键值对
 * 
 */
#[derive(Debug, Clone)]
pub struct KVPairStringString {
    /**
     * 键
     * 
     */
    pub key: Option<String>,
    /**
     * 值
     * 
     */
    pub value: Option<String>,
    
}




/**
 * 数据标识
 * 
 */
#[derive(Debug, Clone)]
pub struct IdentityPairStringString {
    /**
     * 标识
     * 
     */
    pub id: Option<String>,
    /**
     * 值
     * 
     */
    pub value: Option<String>,
    
}




/**
 * 数据标识
 * 
 */
#[derive(Debug, Clone)]
pub struct IdentityPairLongString {
    /**
     * 标识
     * 
     */
    pub id: Option<u64>,
    /**
     * 值
     * 
     */
    pub value: Option<String>,
    
}




/**
 * 数据标识
 * 
 */
#[derive(Debug, Clone)]
pub struct IdentityPairLongLong {
    /**
     * 标识
     * 
     */
    pub id: Option<u64>,
    /**
     * 值
     * 
     */
    pub value: Option<u64>,
    
}




/**
 * 推广表单
 * 
 */
#[derive(Debug, Clone)]
pub struct GlobalPromotionForm {
    /**
     * 推广ID
     * 
     */
    pub adId: Option<u64>,
    /**
     * 答案
     * 
     */
    pub answer: Option<String>,
    /**
     * 推广任务ID
     * 
     */
    pub jobId: Option<u64>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderState {
    /**
     * 
     * 
     */
    pub amount: Option<usize>,
    /**
     * 
     * 
     */
    pub orderId: Option<u64>,
    /**
     * 
     * 
     */
    pub orderState: Option<usize>,
    /**
     * 
     * 
     */
    pub orderStateRemark: Option<String>,
    
}




/**
 * 
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderDiscountPreview {
    /**
     * 优惠券集
     * 
     */
    pub couponIds: Option<Vec<usize>>,
    /**
     * 产品编码
     * 
     */
    pub productNo: Option<String>,
    /**
     * 产品类型
     * 
     */
    pub productType: Option<String>,
    
}




/**
 * 创建订单
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderCreate {
    /**
     * 抵扣优惠券
     * 
     */
    pub couponIds: Option<Vec<usize>>,
    /**
     * 抵扣金额
     * 
     */
    pub discountAmount: Option<usize>,
    /**
     * 订单金额
     * 
     */
    pub orderAmount: Option<usize>,
    /**
     * 订单备注
     * 
     */
    pub orderRemark: Option<String>,
    /**
     * 订单类型
     * 
     */
    pub orderType: Option<String>,
    /**
     * 支付方式
     * 
     */
    pub paymentType: Option<String>,
    /**
     * 产品编码
     * 
     */
    pub productNo: Option<String>,
    /**
     * 产品类型
     * 
     */
    pub productType: Option<String>,
    /**
     * 客户端请求ID
     * 
     */
    pub reqId: Option<String>,
    
}




/**
 * 取消订单
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderCancel {
    /**
     * 类型
     * 
     */
    pub kind: Option<String>,
    /**
     * 订单金额
     * 
     */
    pub orderAmount: Option<u64>,
    /**
     * 订单ID
     * 
     */
    pub orderId: Option<u64>,
    /**
     * 理由
     * 
     */
    pub reason: Option<String>,
    
}




/**
 * OSS文件地址
 * 
 */
#[derive(Debug, Clone)]
pub struct FileLocation {
    /**
     * 文件上传格式
     * 
     */
    pub contentType: Option<String>,
    /**
     * 目录
     * 
     */
    pub directory: Option<String>,
    /**
     * 文件格式
     * 
     */
    pub extension: Option<String>,
    /**
     * 权限标识
     * 
     */
    pub permit: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultLong {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<u64>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultMapStringString {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Map<String,String>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 设备激活
 * 
 */
#[derive(Debug, Clone)]
pub struct DeviceActivation {
    /**
     * 
     * 
     */
    pub caid: Option<String>,
    /**
     * 
     * 
     */
    pub idfa: Option<String>,
    /**
     * 
     * 
     */
    pub imei: Option<String>,
    /**
     * 
     * 
     */
    pub mac: Option<String>,
    /**
     * 
     * 
     */
    pub oaid: Option<String>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioPackageAuditResult {
    /**
     * 
     * 
     */
    pub auditFlag: Option<usize>,
    /**
     * 
     * 
     */
    pub auditRemark: Option<String>,
    /**
     * 
     * 
     */
    pub auditTime: Option<String>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct AdBannerViewEntity {
    /**
     * 
     * 
     */
    pub image: Option<String>,
    /**
     * 
     * 
     */
    pub jumpType: Option<String>,
    /**
     * 
     * 
     */
    pub jumpUrl: Option<String>,
    /**
     * 
     * 
     */
    pub name: Option<String>,
    
}




/**
 * 语音文件
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioFileEntity {
    /**
     * 播放音频文件地址
     * 
     */
    pub audio: Option<String>,
    /**
     * 
     * 
     */
    pub audioArgs: Option<JsonNode>,
    /**
     * 文件内容
     * 
     */
    pub content: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 原始音频文件地址
     * 
     */
    pub file: Option<String>,
    /**
     * 是否隐藏
     * 
     */
    pub hiddenFlag: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 用户ID
     * 
     */
    pub ownerId: Option<String>,
    /**
     * 语音包ID
     * 
     */
    pub packageId: Option<u64>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 解锁方式
     * 
     */
    pub unlockType: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultGlobalPromotionFinish {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<GlobalPromotionFinish>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataUserInviteLog {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataUserInviteLog>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultRewardTip {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<RewardTip>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultStatisticReference {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<StatisticReference>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUpgradeResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UpgradeResult>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUserBehaviorReference {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UserBehaviorReference>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUserFinanceAccount {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UserFinanceAccount>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUserGrowupState {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UserGrowupState>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUserPrincipalEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UserPrincipalEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUserVipRechargeAccount {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<UserVipRechargeAccount>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultVipProductEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<VipProductEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataUserCouponEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataUserCouponEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataCashOrderEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataCashOrderEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioFileEntityViewEntry {
    /**
     * 播放音频文件地址
     * 
     */
    pub audio: Option<String>,
    /**
     * 
     * 
     */
    pub audioArgs: Option<JsonNode>,
    /**
     * 文件内容
     * 
     */
    pub content: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 
     * 
     */
    pub favorite: Option<bool>,
    /**
     * 原始音频文件地址
     * 
     */
    pub file: Option<String>,
    /**
     * 是否隐藏
     * 
     */
    pub hiddenFlag: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 用户ID
     * 
     */
    pub ownerId: Option<String>,
    /**
     * 语音包ID
     * 
     */
    pub packageId: Option<u64>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 解锁方式
     * 
     */
    pub unlockType: Option<String>,
    /**
     * 
     * 
     */
    pub unlocked: Option<bool>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataAudioPackageEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataAudioPackageEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 广告推广任务
 * 
 */
#[derive(Debug, Clone)]
pub struct AdPromotionJobEntity {
    /**
     * 
     * 
     */
    pub answers: Option<JsonNode>,
    /**
     * 已完成数
     * 
     */
    pub completedQuantity: Option<u64>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 是否启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 任务名称
     * 
     */
    pub jobName: Option<String>,
    /**
     * 任务类型
     * 
     */
    pub jobType: Option<String>,
    /**
     * 跳转图片
     * 
     */
    pub jumpImage: Option<String>,
    /**
     * 跳转类型
     * 
     */
    pub jumpType: Option<String>,
    /**
     * 跳转地址,image,word 时有效
     * 
     */
    pub jumpUrl: Option<String>,
    /**
     * 广告推广ID
     * 
     */
    pub promotionId: Option<u64>,
    /**
     * 任务数量
     * 
     */
    pub quantity: Option<u64>,
    /**
     * 任务补充数量
     * 
     */
    pub repairQuantity: Option<u64>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultAdInterestResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<AdInterestResult>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 提现申请订单
 * 
 */
#[derive(Debug, Clone)]
pub struct CashOrderCreate {
    /**
     * 提现金额
     * 
     */
    pub amount: Option<usize>,
    /**
     * 备注
     * 
     */
    pub remark: Option<String>,
    /**
     * 凭证图片
     * 
     */
    pub voucherImage: Option<String>,
    /**
     * 凭证名称
     * 
     */
    pub voucherName: Option<String>,
    /**
     * 凭证帐号
     * 
     */
    pub voucherNo: Option<String>,
    /**
     * 
     * 
     */
    pub voucherPayload: Option<JsonNode>,
    /**
     * 凭证类型
     * 
     */
    pub voucherType: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultAudioPackageAuditResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<AudioPackageAuditResult>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataAudioFileEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataAudioFileEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListVipProductEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<VipProductEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListUserCouponEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<UserCouponEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListTimbreProductEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<TimbreProductEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 创建订单结果
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderCreateResult {
    /**
     * 支付金额
     * 
     */
    pub amount: Option<usize>,
    /**
     * 订单过期时长(分钟)
     * 
     */
    pub expireDuration: Option<usize>,
    /**
     * 订单过期时间
     * 
     */
    pub expireTime: Option<String>,
    /**
     * 订单ID
     * 
     */
    pub orderId: Option<u64>,
    /**
     * 
     * 
     */
    pub payload: Option<JsonNode>,
    /**
     * 支付ID
     * 
     */
    pub paymentId: Option<u64>,
    /**
     * 支付类型
     * 
     */
    pub paymentType: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListKVPairStringString {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<KVPairStringString>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct GeneralOrderDiscountPreviewResult {
    /**
     * 累计折扣金额
     * 
     */
    pub discountAmount: Option<usize>,
    /**
     * 预选抵扣信息
     * 
     */
    pub discounts: Option<Vec<UserCouponEntity>>,
    /**
     * 订单金额
     * 
     */
    pub orderAmount: Option<usize>,
    /**
     * 预计支付金额
     * 
     */
    pub paymentAmount: Option<usize>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListIdentityPairStringString {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<IdentityPairStringString>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct GlobalPromotionFinish {
    /**
     * 是否结束
     * 
     */
    pub finished: Option<bool>,
    /**
     * 
     * 
     */
    pub rewardTip: Option<RewardTip>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAudioPackageEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AudioPackageEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultAudioPackageViewEntry {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<AudioPackageViewEntry>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAudioFileEntityViewEntry {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AudioFileEntityViewEntry>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAudioFileEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AudioFileEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAdPromotionJobEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AdPromotionJobEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAdNextDayLaunchResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AdNextDayLaunchResult>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListAdBannerViewEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<AdBannerViewEntity>>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataAudioFileEntity {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<AudioFileEntity>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataAudioPackageEntity {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<AudioPackageEntity>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataCashOrderEntity {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<CashOrderEntity>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataUserCouponEntity {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<UserCouponEntity>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataUserFinanceEntryDetail {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<UserFinanceEntryDetail>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataUserInviteLog {
    /**
     * 数据集
     * 
     */
    pub entities: Option<Vec<UserInviteLog>>,
    /**
     * 当前页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 每页数量
     * 
     */
    pub size: Option<usize>,
    /**
     * 总条数
     * 
     */
    pub totalCount: Option<usize>,
    /**
     * 总页数
     * 
     */
    pub totalPages: Option<usize>,
    
}




/**
 * 举报内容
 * 
 */
#[derive(Debug, Clone)]
pub struct ReportContentEntity {
    /**
     * 内容
     * 
     */
    pub content: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 反馈
     * 
     */
    pub feedback: Option<String>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 
     * 
     */
    pub payload: Option<JsonNode>,
    /**
     * 理由
     * 
     */
    pub reason: Option<String>,
    /**
     * 资源ID
     * 
     */
    pub resourceId: Option<String>,
    /**
     * 资源类型
     * 
     */
    pub resourceType: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultKVPairStringString {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<KVPairStringString>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultGlobalPromotionProgress {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<GlobalPromotionProgress>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 音色产品
 * 
 */
#[derive(Debug, Clone)]
pub struct TimbreProductEntity {
    /**
     * 试听文件
     * 
     */
    pub audio: Option<String>,
    /**
     * 音色编码
     * 
     */
    pub code: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 是否启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 音色性别
     * 
     */
    pub gender: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 图片
     * 
     */
    pub image: Option<String>,
    /**
     * 服务商
     * 
     */
    pub isp: Option<String>,
    /**
     * 服务商参数
     * 
     */
    pub ispParameter: Option<String>,
    /**
     * 音色名称
     * 
     */
    pub name: Option<String>,
    /**
     * 
     * 
     */
    pub recommendTag: Option<JsonNode>,
    /**
     * 场景
     * 
     */
    pub scene: Option<String>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 解锁类型
     * 
     */
    pub unlockType: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataUserFinanceEntryDetail {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataUserFinanceEntryDetail>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 用户语音包上传表单
 * 
 */
#[derive(Debug, Clone)]
pub struct UserAudioPackageUploadForm {
    /**
     * 音频文件
     * 
     */
    pub files: Option<Vec<KVPairStringString>>,
    /**
     * ID
     * 
     */
    pub id: Option<u64>,
    /**
     * 封面图
     * 
     */
    pub image: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 单价
     * 
     */
    pub price: Option<usize>,
    
}




/**
 * 
 * 
 */
#[derive(Debug, Clone)]
pub struct UserBehaviorArg {
    /**
     * 行为事件
     * 
     */
    pub action: Option<String>,
    /**
     * 
     * 
     */
    pub actionPayload: Option<JsonNode>,
    /**
     * 资源ID
     * 
     */
    pub mainstayId: Option<String>,
    /**
     * 资源作者ID
     * 
     */
    pub mainstayOwnerBy: Option<String>,
    /**
     * 资源类型
     * 
     */
    pub mainstayType: Option<String>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultGeneralOrderState {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<GeneralOrderState>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 用户优惠券
 * 
 */
#[derive(Debug, Clone)]
pub struct UserCouponEntity {
    /**
     * 优惠券编码
     * 
     */
    pub couponCode: Option<String>,
    /**
     * 优惠券ID
     * 
     */
    pub couponId: Option<u64>,
    /**
     * 优惠券名称
     * 
     */
    pub couponName: Option<String>,
    /**
     * 优惠券类型
     * 
     */
    pub couponType: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 优惠券生效期
     * 
     */
    pub effectTime: Option<String>,
    /**
     * 优惠券有效期
     * 
     */
    pub expireTime: Option<String>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 订单ID
     * 
     */
    pub orderId: Option<u64>,
    /**
     * 来源
     * 
     */
    pub origin: Option<String>,
    /**
     * 来源标识
     * 
     */
    pub originId: Option<String>,
    /**
     * 
     * 
     */
    pub snapshot: Option<JsonNode>,
    /**
     * 状态
     * 
     */
    pub state: Option<usize>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    /**
     * 用户Id
     * 
     */
    pub userId: Option<u64>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultGeneralOrderDiscountPreviewResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<GeneralOrderDiscountPreviewResult>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultGeneralOrderCreateResult {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<GeneralOrderCreateResult>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 业务响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultCashOrderEntity {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 响应时间戳
     * 
     */
    pub currentTime: Option<u64>,
    /**
     * 
     * 
     */
    pub data: Option<CashOrderEntity>,
    /**
     * 业务异常信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    /**
     * 成功
     * 
     */
    pub status: Option<bool>,
    
}




/**
 * 提现订单
 * 
 */
#[derive(Debug, Clone)]
pub struct CashOrderEntity {
    /**
     * 申请金额
     * 
     */
    pub applyAmount: Option<usize>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 手续费
     * 
     */
    pub feeAmount: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 操作员ID
     * 
     */
    pub operateId: Option<String>,
    /**
     * 打款金额
     * 
     */
    pub paymentAmount: Option<usize>,
    /**
     * 打款ID
     * 
     */
    pub paymentId: Option<u64>,
    /**
     * 打款凭证
     * 
     */
    pub paymentImage: Option<String>,
    /**
     * 
     * 
     */
    pub paymentPayload: Option<JsonNode>,
    /**
     * 打款时间
     * 
     */
    pub paymentTime: Option<String>,
    /**
     * 收款金额
     * 
     */
    pub receiptAmount: Option<usize>,
    /**
     * 备注
     * 
     */
    pub remark: Option<String>,
    /**
     * 状态
     * 
     */
    pub state: Option<usize>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    /**
     * 申请人ID
     * 
     */
    pub userId: Option<u64>,
    /**
     * 凭证图片
     * 
     */
    pub voucherImage: Option<String>,
    /**
     * 凭证名称
     * 
     */
    pub voucherName: Option<String>,
    /**
     * 凭证帐号
     * 
     */
    pub voucherNo: Option<String>,
    /**
     * 
     * 
     */
    pub voucherPayload: Option<JsonNode>,
    /**
     * 凭证类型
     * 
     */
    pub voucherType: Option<String>,
    
}




/**
 * 用户主体
 * 
 */
#[derive(Debug, Clone)]
pub struct UserPrincipalEntity {
    /**
     * 地址
     * 
     */
    pub address: Option<String>,
    /**
     * 
     * 
     */
    pub addressPca: Option<JsonNode>,
    /**
     * 头像
     * 
     */
    pub avatar: Option<String>,
    /**
     * 生日
     * 
     */
    pub birthday: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 邮箱
     * 
     */
    pub email: Option<String>,
    /**
     * 状态
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 性别
     * 
     */
    pub gender: Option<usize>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 邀请码（别名）
     * 
     */
    pub inviteAlias: Option<String>,
    /**
     * 邀请码
     * 
     */
    pub inviteCode: Option<String>,
    /**
     * 登陆账号
     * 
     */
    pub loginId: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 昵称
     * 
     */
    pub nickName: Option<String>,
    /**
     * 
     * 
     */
    pub profile: Option<JsonNode>,
    /**
     * 推荐人-用户ID
     * 
     */
    pub referrerId: Option<u64>,
    /**
     * 推荐人-邀请码
     * 
     */
    pub referrerInviteCode: Option<String>,
    /**
     * 是否为重新注册
     * 
     */
    pub resignFlag: Option<usize>,
    /**
     * 类型
     * 
     */
    pub type: Option<String>,
    /**
     * 唯一手机号
     * 
     */
    pub uniquePhone: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 推广进度
 * 
 */
#[derive(Debug, Clone)]
pub struct GlobalPromotionProgress {
    /**
     * 
     * 
     */
    pub ad: Option<AdPromotionEntity>,
    /**
     * 是否结束
     * 
     */
    pub finished: Option<bool>,
    /**
     * 当前任务序号
     * 
     */
    pub index: Option<usize>,
    /**
     * 
     * 
     */
    pub job: Option<AdPromotionJobEntity>,
    /**
     * 
     * 
     */
    pub rewardTip: Option<RewardTip>,
    
}




/**
 * 语音包
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioPackageEntity {
    /**
     * 审核状态
     * 
     */
    pub auditFlag: Option<usize>,
    /**
     * 审核备注
     * 
     */
    pub auditRemark: Option<String>,
    /**
     * 审核时间
     * 
     */
    pub auditTime: Option<String>,
    /**
     * 分类ID
     * 
     */
    pub categoryId: Option<String>,
    /**
     * 分类名称
     * 
     */
    pub categoryName: Option<String>,
    /**
     * 
     * 
     */
    pub categoryTag: Option<JsonNode>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 备注
     * 
     */
    pub description: Option<String>,
    /**
     * 启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 
     * 
     */
    pub fakeTag: Option<JsonNode>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 封面图
     * 
     */
    pub image: Option<String>,
    /**
     * 所属类型
     * 
     */
    pub mode: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 所属用户ID
     * 
     */
    pub ownerId: Option<String>,
    /**
     * 单价
     * 
     */
    pub price: Option<usize>,
    /**
     * 
     * 
     */
    pub recommendTag: Option<JsonNode>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 解锁方式
     * 
     */
    pub unlockType: Option<String>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    
}




/**
 * 数据
 * 
 */
#[derive(Debug, Clone)]
pub struct AudioPackageViewEntry {
    /**
     * 审核状态
     * 
     */
    pub auditFlag: Option<usize>,
    /**
     * 审核备注
     * 
     */
    pub auditRemark: Option<String>,
    /**
     * 审核时间
     * 
     */
    pub auditTime: Option<String>,
    /**
     * 分类ID
     * 
     */
    pub categoryId: Option<String>,
    /**
     * 分类名称
     * 
     */
    pub categoryName: Option<String>,
    /**
     * 
     * 
     */
    pub categoryTag: Option<JsonNode>,
    /**
     * 
     * 
     */
    pub coupons: Option<Vec<UserCouponEntity>>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 备注
     * 
     */
    pub description: Option<String>,
    /**
     * 启用
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 
     * 
     */
    pub fakeTag: Option<JsonNode>,
    /**
     * 
     * 
     */
    pub freeCount: Option<u64>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 封面图
     * 
     */
    pub image: Option<String>,
    /**
     * 所属类型
     * 
     */
    pub mode: Option<String>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 所属用户ID
     * 
     */
    pub ownerId: Option<String>,
    /**
     * 单价
     * 
     */
    pub price: Option<usize>,
    /**
     * 
     * 
     */
    pub recommendTag: Option<JsonNode>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 
     * 
     */
    pub totalCount: Option<u64>,
    /**
     * 解锁方式
     * 
     */
    pub unlockType: Option<String>,
    /**
     * 
     * 
     */
    pub unlocked: Option<bool>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    /**
     * 
     * 
     */
    pub vipCount: Option<u64>,
    
}




/**
 * 广告推广
 * 
 */
#[derive(Debug, Clone)]
pub struct AdPromotionEntity {
    /**
     * 开始时间
     * 
     */
    pub beginTime: Option<String>,
    /**
     * 渠道
     * 
     */
    pub channel: Option<String>,
    /**
     * 创建时间
     * 
     */
    pub createdAt: Option<String>,
    /**
     * 备注
     * 
     */
    pub description: Option<String>,
    /**
     * 启用标识
     * 
     */
    pub enableFlag: Option<usize>,
    /**
     * 结束时间
     * 
     */
    pub endTime: Option<String>,
    /**
     * 渠道应用icon
     * 
     */
    pub icon: Option<String>,
    /**
     * 主键
     * 
     */
    pub id: Option<u64>,
    /**
     * 图片
     * 
     */
    pub image: Option<String>,
    /**
     * 任务图片为主
     * 
     */
    pub imagePosition: Option<usize>,
    /**
     * 
     * 
     */
    pub installs: Option<JsonNode>,
    /**
     * 跳转地址
     * 
     */
    pub jumpUrl: Option<String>,
    /**
     * 唤起标识
     * 
     */
    pub launchFlag: Option<usize>,
    /**
     * 唤起地址
     * 
     */
    pub launchUrl: Option<String>,
    /**
     * 
     * 
     */
    pub mistakeAnswers: Option<JsonNode>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 
     * 
     */
    pub rewards: Option<JsonNode>,
    /**
     * 
     * 
     */
    pub rightAnswers: Option<JsonNode>,
    /**
     * 风控时间(秒)
     * 
     */
    pub riskInterval: Option<usize>,
    /**
     * 排序
     * 
     */
    pub sort: Option<usize>,
    /**
     * 
     * 
     */
    pub tips: Option<JsonNode>,
    /**
     * 解锁任务阈值
     * 
     */
    pub unlockThreshold: Option<usize>,
    /**
     * 解锁类型
     * 
     */
    pub unlockType: Option<usize>,
    /**
     * 更新时间
     * 
     */
    pub updatedAt: Option<String>,
    /**
     * 视频教程
     * 
     */
    pub video: Option<String>,
    
}




/**
 * 兴趣社
 * 
 */
#[derive(Debug, Clone)]
pub struct AdInterestResult {
    /**
     * 语音包收入
     * 
     */
    pub audioIncome: Option<Vec<UserItem>>,
    /**
     * 本月冠军
     * 
     */
    pub champion: Option<Vec<UserItem>>,
    /**
     * 合作伙伴
     * 
     */
    pub partner: Option<Vec<UserItem>>,
    /**
     * 语音包排行
     * 
     */
    pub ranking: Option<Vec<UserItem>>,
    /**
     * 视频收入
     * 
     */
    pub videoIncome: Option<Vec<UserItem>>,
    /**
     * 热门视频
     * 
     */
    pub videos: Option<Vec<VideoItem>>,
    
}




