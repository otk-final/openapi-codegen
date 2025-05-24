




/**
 * 
 * 
 */
public class AudioCouponUnlockForm
{
    /**
     * 
     * 
     */
    public long? CouponId { get; set; }
    /**
     * 
     * 
     */
    public long? PackageId { get; set; }
    
}




/**
 * 会员产品
 * 
 */
public class VipProductEntity
{
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 备注
     * 
     */
    public string? Description { get; set; }
    /**
     * 时长
     * 
     */
    public string? Duration { get; set; }
    /**
     * 启用标识
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 封面图
     * 
     */
    public string? Image { get; set; }
    /**
     * 划线价
     * 
     */
    public int? ListPrice { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 编码
     * 
     */
    public string? No { get; set; }
    /**
     * 售价
     * 
     */
    public int? Price { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 标题
     * 
     */
    public string? Title { get; set; }
    /**
     * 可用时长（秒）
     * 
     */
    public long? Token { get; set; }
    /**
     * 类型
     * 
     */
    public string? Type { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 数据
 * 
 */
public class AdNextDayLaunchResult
{
    /**
     * 
     * 
     */
    public string? Icon { get; set; }
    /**
     * 标识
     * 
     */
    public long? Id { get; set; }
    /**
     * 
     * 
     */
    public string? JumpUrl { get; set; }
    /**
     * 
     * 
     */
    public string? LaunchUrl { get; set; }
    /**
     * 
     * 
     */
    public string? Name { get; set; }
    /**
     * 值
     * 
     */
    public string? Value { get; set; }
    
}




/**
 * 视频
 * 
 */
public class VideoItem
{
    /**
     * 封面图
     * 
     */
    public string? Image { get; set; }
    /**
     * 跳转地址
     * 
     */
    public string? JumpUrl { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 视频地址
     * 
     */
    public string? Video { get; set; }
    
}




/**
 * 数据
 * 
 */
public class UserVipRechargeAccount
{
    /**
     * 可用语音包数量
     * 
     */
    public string? AvailableAudioCount { get; set; }
    /**
     * 可用时长(秒)
     * 
     */
    public long? AvailableToken { get; set; }
    /**
     * 到期时间
     * 
     */
    public string? ExpireTime { get; set; }
    /**
     * 
     * 
     */
    public int? TotalAmount { get; set; }
    /**
     * 
     * 
     */
    public long? TotalCount { get; set; }
    /**
     * 状态
     * 
     */
    public int? Valid { get; set; }
    /**
     * 状态文案
     * 
     */
    public string? ValidDescription { get; set; }
    
}




/**
 * 用户
 * 
 */
public class UserItem
{
    /**
     * 头像地址
     * 
     */
    public string? Avatar { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 下载量 / 收益
     * 
     */
    public int? Weight { get; set; }
    
}




/**
 * 用户邀请记录
 * 
 */
public class UserInviteLog
{
    /**
     * 邀请人头像
     * 
     */
    public string? Avatar { get; set; }
    /**
     * 邀请时间
     * 
     */
    public string? InviteTime { get; set; }
    /**
     * 邀请人名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 邀请人ID
     * 
     */
    public long? UserId { get; set; }
    
}




/**
 * 用户初始化基础信息
 * 
 */
public class UserInitProfile
{
    /**
     * 头像
     * 
     */
    public string? Avatar { get; set; }
    /**
     * 生日
     * 
     */
    public string? Birthday { get; set; }
    /**
     * 性别
     * 
     */
    public int? Gender { get; set; }
    /**
     * 邀请码
     * 
     */
    public string? InviteCode { get; set; }
    /**
     * 昵称
     * 
     */
    public string? Name { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultBoolean
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public bool? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 用户成长状态
 * 
 */
public class UserGrowupState
{
    /**
     * 当前经验值
     * 
     */
    public int? CurrentEmpiric { get; set; }
    /**
     * 等级
     * 
     */
    public int? Level { get; set; }
    /**
     * 注册天数
     * 
     */
    public int? RegisterDay { get; set; }
    /**
     * 累计经验值
     * 
     */
    public int? TotalEmpiric { get; set; }
    
}




/**
 * 用户资金明细
 * 
 */
public class UserFinanceEntryDetail
{
    /**
     * 交易金额
     * 
     */
    public int? TransAmount { get; set; }
    /**
     * 交易备注
     * 
     */
    public string? TransRemark { get; set; }
    /**
     * 交易时间
     * 
     */
    public string? TransTime { get; set; }
    /**
     * 交易类型
     * 
     */
    public string? TransType { get; set; }
    /**
     * 用户头像
     * 
     */
    public string? UserAvatar { get; set; }
    /**
     * 用户ID
     * 
     */
    public long? UserId { get; set; }
    /**
     * 用户昵称
     * 
     */
    public string? UserName { get; set; }
    
}




/**
 * 用户账户信息
 * 
 */
public class UserFinanceAccount
{
    /**
     * 当前申请提现
     * 
     */
    public int? ApplyCash { get; set; }
    /**
     * 语音包奖励
     * 
     */
    public int? AudioReward { get; set; }
    /**
     * 余额
     * 
     */
    public int? Balance { get; set; }
    /**
     * 任务奖励
     * 
     */
    public int? JobReward { get; set; }
    /**
     * 平台奖励
     * 
     */
    public int? PlatformReward { get; set; }
    /**
     * 累计提现
     * 
     */
    public int? TotalCash { get; set; }
    /**
     * 累计收益
     * 
     */
    public int? TotalProfit { get; set; }
    /**
     * 会员奖励
     * 
     */
    public int? VipReward { get; set; }
    
}




/**
 * 行为统计
 * 
 */
public class UserBehaviorReference
{
    /**
     * 评论
     * 
     */
    public bool? Comment { get; set; }
    /**
     * 踩
     * 
     */
    public bool? Dislike { get; set; }
    /**
     * 收藏
     * 
     */
    public bool? Favorite { get; set; }
    /**
     * 点赞
     * 
     */
    public bool? Like { get; set; }
    /**
     * 分享
     * 
     */
    public bool? Share { get; set; }
    /**
     * 浏览
     * 
     */
    public bool? View { get; set; }
    
}




/**
 * 升级
 * 
 */
public class UpgradeResult
{
    /**
     * 是否弹窗
     * 
     */
    public bool? Alter { get; set; }
    /**
     * 更新说明
     * 
     */
    public string? Description { get; set; }
    /**
     * 下载地址
     * 
     */
    public string? Download { get; set; }
    /**
     * 是否强制升级
     * 
     */
    public bool? Focus { get; set; }
    /**
     * 当前版本
     * 
     */
    public string? Version { get; set; }
    
}




/**
 * 数量统计
 * 
 */
public class StatisticReference
{
    /**
     * 评论数
     * 
     */
    public long? CommentCount { get; set; }
    /**
     * 不喜欢数
     * 
     */
    public long? DislikeCount { get; set; }
    /**
     * 收藏数
     * 
     */
    public long? FavoriteCount { get; set; }
    /**
     * 点赞数
     * 
     */
    public long? LikeCount { get; set; }
    /**
     * 分享数
     * 
     */
    public long? ShareCount { get; set; }
    /**
     * 浏览数
     * 
     */
    public long? ViewCount { get; set; }
    
}




/**
 * 奖励提示
 * 
 */
public class RewardTip
{
    /**
     * 
     * 
     */
    public string? Content { get; set; }
    /**
     * 
     * 
     */
    public int? Progress { get; set; }
    /**
     * 
     * 
     */
    public string? TotalContent { get; set; }
    
}




/**
 * 键值对
 * 
 */
public class KVPairStringString
{
    /**
     * 键
     * 
     */
    public string? Key { get; set; }
    /**
     * 值
     * 
     */
    public string? Value { get; set; }
    
}




/**
 * 数据标识
 * 
 */
public class IdentityPairStringString
{
    /**
     * 标识
     * 
     */
    public string? Id { get; set; }
    /**
     * 值
     * 
     */
    public string? Value { get; set; }
    
}




/**
 * 数据标识
 * 
 */
public class IdentityPairLongString
{
    /**
     * 标识
     * 
     */
    public long? Id { get; set; }
    /**
     * 值
     * 
     */
    public string? Value { get; set; }
    
}




/**
 * 数据标识
 * 
 */
public class IdentityPairLongLong
{
    /**
     * 标识
     * 
     */
    public long? Id { get; set; }
    /**
     * 值
     * 
     */
    public long? Value { get; set; }
    
}




/**
 * 推广表单
 * 
 */
public class GlobalPromotionForm
{
    /**
     * 推广ID
     * 
     */
    public long? AdId { get; set; }
    /**
     * 答案
     * 
     */
    public string? Answer { get; set; }
    /**
     * 推广任务ID
     * 
     */
    public long? JobId { get; set; }
    
}




/**
 * 数据
 * 
 */
public class GeneralOrderState
{
    /**
     * 
     * 
     */
    public int? Amount { get; set; }
    /**
     * 
     * 
     */
    public long? OrderId { get; set; }
    /**
     * 
     * 
     */
    public int? OrderState { get; set; }
    /**
     * 
     * 
     */
    public string? OrderStateRemark { get; set; }
    
}




/**
 * 
 * 
 */
public class GeneralOrderDiscountPreview
{
    /**
     * 优惠券集
     * 
     */
    public List<int>? CouponIds { get; set; }
    /**
     * 产品编码
     * 
     */
    public string? ProductNo { get; set; }
    /**
     * 产品类型
     * 
     */
    public string? ProductType { get; set; }
    
}




/**
 * 创建订单
 * 
 */
public class GeneralOrderCreate
{
    /**
     * 抵扣优惠券
     * 
     */
    public List<int>? CouponIds { get; set; }
    /**
     * 抵扣金额
     * 
     */
    public int? DiscountAmount { get; set; }
    /**
     * 订单金额
     * 
     */
    public int? OrderAmount { get; set; }
    /**
     * 订单备注
     * 
     */
    public string? OrderRemark { get; set; }
    /**
     * 订单类型
     * 
     */
    public string? OrderType { get; set; }
    /**
     * 支付方式
     * 
     */
    public string? PaymentType { get; set; }
    /**
     * 产品编码
     * 
     */
    public string? ProductNo { get; set; }
    /**
     * 产品类型
     * 
     */
    public string? ProductType { get; set; }
    /**
     * 客户端请求ID
     * 
     */
    public string? ReqId { get; set; }
    
}




/**
 * 取消订单
 * 
 */
public class GeneralOrderCancel
{
    /**
     * 类型
     * 
     */
    public string? Kind { get; set; }
    /**
     * 订单金额
     * 
     */
    public long? OrderAmount { get; set; }
    /**
     * 订单ID
     * 
     */
    public long? OrderId { get; set; }
    /**
     * 理由
     * 
     */
    public string? Reason { get; set; }
    
}




/**
 * OSS文件地址
 * 
 */
public class FileLocation
{
    /**
     * 文件上传格式
     * 
     */
    public string? ContentType { get; set; }
    /**
     * 目录
     * 
     */
    public string? Directory { get; set; }
    /**
     * 文件格式
     * 
     */
    public string? Extension { get; set; }
    /**
     * 权限标识
     * 
     */
    public string? Permit { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultLong
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public long? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultMapStringString
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public Dictionary<string,string>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 设备激活
 * 
 */
public class DeviceActivation
{
    /**
     * 
     * 
     */
    public string? Caid { get; set; }
    /**
     * 
     * 
     */
    public string? Idfa { get; set; }
    /**
     * 
     * 
     */
    public string? Imei { get; set; }
    /**
     * 
     * 
     */
    public string? Mac { get; set; }
    /**
     * 
     * 
     */
    public string? Oaid { get; set; }
    
}




/**
 * 数据
 * 
 */
public class AudioPackageAuditResult
{
    /**
     * 
     * 
     */
    public int? AuditFlag { get; set; }
    /**
     * 
     * 
     */
    public string? AuditRemark { get; set; }
    /**
     * 
     * 
     */
    public string? AuditTime { get; set; }
    
}




/**
 * 数据
 * 
 */
public class AdBannerViewEntity
{
    /**
     * 
     * 
     */
    public string? Image { get; set; }
    /**
     * 
     * 
     */
    public string? JumpType { get; set; }
    /**
     * 
     * 
     */
    public string? JumpUrl { get; set; }
    /**
     * 
     * 
     */
    public string? Name { get; set; }
    
}




/**
 * 语音文件
 * 
 */
public class AudioFileEntity
{
    /**
     * 播放音频文件地址
     * 
     */
    public string? Audio { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? AudioArgs { get; set; }
    /**
     * 文件内容
     * 
     */
    public string? Content { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 原始音频文件地址
     * 
     */
    public string? File { get; set; }
    /**
     * 是否隐藏
     * 
     */
    public int? HiddenFlag { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 用户ID
     * 
     */
    public string? OwnerId { get; set; }
    /**
     * 语音包ID
     * 
     */
    public long? PackageId { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 解锁方式
     * 
     */
    public string? UnlockType { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultGlobalPromotionFinish
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public GlobalPromotionFinish? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataUserInviteLog
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataUserInviteLog? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultRewardTip
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public RewardTip? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultStatisticReference
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public StatisticReference? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUpgradeResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UpgradeResult? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUserBehaviorReference
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UserBehaviorReference? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUserFinanceAccount
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UserFinanceAccount? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUserGrowupState
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UserGrowupState? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUserPrincipalEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UserPrincipalEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultUserVipRechargeAccount
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public UserVipRechargeAccount? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultVipProductEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public VipProductEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataUserCouponEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataUserCouponEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataCashOrderEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataCashOrderEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 数据
 * 
 */
public class AudioFileEntityViewEntry
{
    /**
     * 播放音频文件地址
     * 
     */
    public string? Audio { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? AudioArgs { get; set; }
    /**
     * 文件内容
     * 
     */
    public string? Content { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 
     * 
     */
    public bool? Favorite { get; set; }
    /**
     * 原始音频文件地址
     * 
     */
    public string? File { get; set; }
    /**
     * 是否隐藏
     * 
     */
    public int? HiddenFlag { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 用户ID
     * 
     */
    public string? OwnerId { get; set; }
    /**
     * 语音包ID
     * 
     */
    public long? PackageId { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 解锁方式
     * 
     */
    public string? UnlockType { get; set; }
    /**
     * 
     * 
     */
    public bool? Unlocked { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataAudioPackageEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataAudioPackageEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 广告推广任务
 * 
 */
public class AdPromotionJobEntity
{
    /**
     * 
     * 
     */
    public JsonNode? Answers { get; set; }
    /**
     * 已完成数
     * 
     */
    public long? CompletedQuantity { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 是否启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 任务名称
     * 
     */
    public string? JobName { get; set; }
    /**
     * 任务类型
     * 
     */
    public string? JobType { get; set; }
    /**
     * 跳转图片
     * 
     */
    public string? JumpImage { get; set; }
    /**
     * 跳转类型
     * 
     */
    public string? JumpType { get; set; }
    /**
     * 跳转地址,image,word 时有效
     * 
     */
    public string? JumpUrl { get; set; }
    /**
     * 广告推广ID
     * 
     */
    public long? PromotionId { get; set; }
    /**
     * 任务数量
     * 
     */
    public long? Quantity { get; set; }
    /**
     * 任务补充数量
     * 
     */
    public long? RepairQuantity { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultAdInterestResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public AdInterestResult? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 提现申请订单
 * 
 */
public class CashOrderCreate
{
    /**
     * 提现金额
     * 
     */
    public int? Amount { get; set; }
    /**
     * 备注
     * 
     */
    public string? Remark { get; set; }
    /**
     * 凭证图片
     * 
     */
    public string? VoucherImage { get; set; }
    /**
     * 凭证名称
     * 
     */
    public string? VoucherName { get; set; }
    /**
     * 凭证帐号
     * 
     */
    public string? VoucherNo { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? VoucherPayload { get; set; }
    /**
     * 凭证类型
     * 
     */
    public string? VoucherType { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultAudioPackageAuditResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public AudioPackageAuditResult? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataAudioFileEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataAudioFileEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListVipProductEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<VipProductEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListUserCouponEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<UserCouponEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListTimbreProductEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<TimbreProductEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 创建订单结果
 * 
 */
public class GeneralOrderCreateResult
{
    /**
     * 支付金额
     * 
     */
    public int? Amount { get; set; }
    /**
     * 订单过期时长(分钟)
     * 
     */
    public int? ExpireDuration { get; set; }
    /**
     * 订单过期时间
     * 
     */
    public string? ExpireTime { get; set; }
    /**
     * 订单ID
     * 
     */
    public long? OrderId { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Payload { get; set; }
    /**
     * 支付ID
     * 
     */
    public long? PaymentId { get; set; }
    /**
     * 支付类型
     * 
     */
    public string? PaymentType { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListKVPairStringString
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<KVPairStringString>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 数据
 * 
 */
public class GeneralOrderDiscountPreviewResult
{
    /**
     * 累计折扣金额
     * 
     */
    public int? DiscountAmount { get; set; }
    /**
     * 预选抵扣信息
     * 
     */
    public List<UserCouponEntity>? Discounts { get; set; }
    /**
     * 订单金额
     * 
     */
    public int? OrderAmount { get; set; }
    /**
     * 预计支付金额
     * 
     */
    public int? PaymentAmount { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListIdentityPairStringString
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<IdentityPairStringString>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 数据
 * 
 */
public class GlobalPromotionFinish
{
    /**
     * 是否结束
     * 
     */
    public bool? Finished { get; set; }
    /**
     * 
     * 
     */
    public RewardTip? RewardTip { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAudioPackageEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AudioPackageEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultAudioPackageViewEntry
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public AudioPackageViewEntry? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAudioFileEntityViewEntry
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AudioFileEntityViewEntry>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAudioFileEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AudioFileEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAdPromotionJobEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AdPromotionJobEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAdNextDayLaunchResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AdNextDayLaunchResult>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultListAdBannerViewEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 数据
     * 
     */
    public List<AdBannerViewEntity>? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataAudioFileEntity
{
    /**
     * 数据集
     * 
     */
    public List<AudioFileEntity>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataAudioPackageEntity
{
    /**
     * 数据集
     * 
     */
    public List<AudioPackageEntity>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataCashOrderEntity
{
    /**
     * 数据集
     * 
     */
    public List<CashOrderEntity>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataUserCouponEntity
{
    /**
     * 数据集
     * 
     */
    public List<UserCouponEntity>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataUserFinanceEntryDetail
{
    /**
     * 数据集
     * 
     */
    public List<UserFinanceEntryDetail>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 分页数据
 * 
 */
public class PageDataUserInviteLog
{
    /**
     * 数据集
     * 
     */
    public List<UserInviteLog>? Entities { get; set; }
    /**
     * 当前页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 每页数量
     * 
     */
    public int? Size { get; set; }
    /**
     * 总条数
     * 
     */
    public int? TotalCount { get; set; }
    /**
     * 总页数
     * 
     */
    public int? TotalPages { get; set; }
    
}




/**
 * 举报内容
 * 
 */
public class ReportContentEntity
{
    /**
     * 内容
     * 
     */
    public string? Content { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 反馈
     * 
     */
    public string? Feedback { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Payload { get; set; }
    /**
     * 理由
     * 
     */
    public string? Reason { get; set; }
    /**
     * 资源ID
     * 
     */
    public string? ResourceId { get; set; }
    /**
     * 资源类型
     * 
     */
    public string? ResourceType { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultKVPairStringString
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public KVPairStringString? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultGlobalPromotionProgress
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public GlobalPromotionProgress? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 音色产品
 * 
 */
public class TimbreProductEntity
{
    /**
     * 试听文件
     * 
     */
    public string? Audio { get; set; }
    /**
     * 音色编码
     * 
     */
    public string? Code { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 是否启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 音色性别
     * 
     */
    public int? Gender { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 图片
     * 
     */
    public string? Image { get; set; }
    /**
     * 服务商
     * 
     */
    public string? Isp { get; set; }
    /**
     * 服务商参数
     * 
     */
    public string? IspParameter { get; set; }
    /**
     * 音色名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? RecommendTag { get; set; }
    /**
     * 场景
     * 
     */
    public string? Scene { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 解锁类型
     * 
     */
    public string? UnlockType { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultPageDataUserFinanceEntryDetail
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public PageDataUserFinanceEntryDetail? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 用户语音包上传表单
 * 
 */
public class UserAudioPackageUploadForm
{
    /**
     * 音频文件
     * 
     */
    public List<KVPairStringString>? Files { get; set; }
    /**
     * ID
     * 
     */
    public long? Id { get; set; }
    /**
     * 封面图
     * 
     */
    public string? Image { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 单价
     * 
     */
    public int? Price { get; set; }
    
}




/**
 * 
 * 
 */
public class UserBehaviorArg
{
    /**
     * 行为事件
     * 
     */
    public string? Action { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? ActionPayload { get; set; }
    /**
     * 资源ID
     * 
     */
    public string? MainstayId { get; set; }
    /**
     * 资源作者ID
     * 
     */
    public string? MainstayOwnerBy { get; set; }
    /**
     * 资源类型
     * 
     */
    public string? MainstayType { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultGeneralOrderState
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public GeneralOrderState? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 用户优惠券
 * 
 */
public class UserCouponEntity
{
    /**
     * 优惠券编码
     * 
     */
    public string? CouponCode { get; set; }
    /**
     * 优惠券ID
     * 
     */
    public long? CouponId { get; set; }
    /**
     * 优惠券名称
     * 
     */
    public string? CouponName { get; set; }
    /**
     * 优惠券类型
     * 
     */
    public string? CouponType { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 优惠券生效期
     * 
     */
    public string? EffectTime { get; set; }
    /**
     * 优惠券有效期
     * 
     */
    public string? ExpireTime { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 订单ID
     * 
     */
    public long? OrderId { get; set; }
    /**
     * 来源
     * 
     */
    public string? Origin { get; set; }
    /**
     * 来源标识
     * 
     */
    public string? OriginId { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Snapshot { get; set; }
    /**
     * 状态
     * 
     */
    public int? State { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    /**
     * 用户Id
     * 
     */
    public long? UserId { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultGeneralOrderDiscountPreviewResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public GeneralOrderDiscountPreviewResult? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultGeneralOrderCreateResult
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public GeneralOrderCreateResult? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 业务响应
 * 
 */
public class ApiResultCashOrderEntity
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 响应时间戳
     * 
     */
    public long? CurrentTime { get; set; }
    /**
     * 
     * 
     */
    public CashOrderEntity? Data { get; set; }
    /**
     * 业务异常信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    /**
     * 成功
     * 
     */
    public bool? Status { get; set; }
    
}




/**
 * 提现订单
 * 
 */
public class CashOrderEntity
{
    /**
     * 申请金额
     * 
     */
    public int? ApplyAmount { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 手续费
     * 
     */
    public int? FeeAmount { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 操作员ID
     * 
     */
    public string? OperateId { get; set; }
    /**
     * 打款金额
     * 
     */
    public int? PaymentAmount { get; set; }
    /**
     * 打款ID
     * 
     */
    public long? PaymentId { get; set; }
    /**
     * 打款凭证
     * 
     */
    public string? PaymentImage { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? PaymentPayload { get; set; }
    /**
     * 打款时间
     * 
     */
    public string? PaymentTime { get; set; }
    /**
     * 收款金额
     * 
     */
    public int? ReceiptAmount { get; set; }
    /**
     * 备注
     * 
     */
    public string? Remark { get; set; }
    /**
     * 状态
     * 
     */
    public int? State { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    /**
     * 申请人ID
     * 
     */
    public long? UserId { get; set; }
    /**
     * 凭证图片
     * 
     */
    public string? VoucherImage { get; set; }
    /**
     * 凭证名称
     * 
     */
    public string? VoucherName { get; set; }
    /**
     * 凭证帐号
     * 
     */
    public string? VoucherNo { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? VoucherPayload { get; set; }
    /**
     * 凭证类型
     * 
     */
    public string? VoucherType { get; set; }
    
}




/**
 * 用户主体
 * 
 */
public class UserPrincipalEntity
{
    /**
     * 地址
     * 
     */
    public string? Address { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? AddressPca { get; set; }
    /**
     * 头像
     * 
     */
    public string? Avatar { get; set; }
    /**
     * 生日
     * 
     */
    public string? Birthday { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 邮箱
     * 
     */
    public string? Email { get; set; }
    /**
     * 状态
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 性别
     * 
     */
    public int? Gender { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 邀请码（别名）
     * 
     */
    public string? InviteAlias { get; set; }
    /**
     * 邀请码
     * 
     */
    public string? InviteCode { get; set; }
    /**
     * 登陆账号
     * 
     */
    public string? LoginId { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 昵称
     * 
     */
    public string? NickName { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Profile { get; set; }
    /**
     * 推荐人-用户ID
     * 
     */
    public long? ReferrerId { get; set; }
    /**
     * 推荐人-邀请码
     * 
     */
    public string? ReferrerInviteCode { get; set; }
    /**
     * 是否为重新注册
     * 
     */
    public int? ResignFlag { get; set; }
    /**
     * 类型
     * 
     */
    public string? Type { get; set; }
    /**
     * 唯一手机号
     * 
     */
    public string? UniquePhone { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 推广进度
 * 
 */
public class GlobalPromotionProgress
{
    /**
     * 
     * 
     */
    public AdPromotionEntity? Ad { get; set; }
    /**
     * 是否结束
     * 
     */
    public bool? Finished { get; set; }
    /**
     * 当前任务序号
     * 
     */
    public int? Index { get; set; }
    /**
     * 
     * 
     */
    public AdPromotionJobEntity? Job { get; set; }
    /**
     * 
     * 
     */
    public RewardTip? RewardTip { get; set; }
    
}




/**
 * 语音包
 * 
 */
public class AudioPackageEntity
{
    /**
     * 审核状态
     * 
     */
    public int? AuditFlag { get; set; }
    /**
     * 审核备注
     * 
     */
    public string? AuditRemark { get; set; }
    /**
     * 审核时间
     * 
     */
    public string? AuditTime { get; set; }
    /**
     * 分类ID
     * 
     */
    public string? CategoryId { get; set; }
    /**
     * 分类名称
     * 
     */
    public string? CategoryName { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? CategoryTag { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 备注
     * 
     */
    public string? Description { get; set; }
    /**
     * 启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? FakeTag { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 封面图
     * 
     */
    public string? Image { get; set; }
    /**
     * 所属类型
     * 
     */
    public string? Mode { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 所属用户ID
     * 
     */
    public string? OwnerId { get; set; }
    /**
     * 单价
     * 
     */
    public int? Price { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? RecommendTag { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 解锁方式
     * 
     */
    public string? UnlockType { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    
}




/**
 * 数据
 * 
 */
public class AudioPackageViewEntry
{
    /**
     * 审核状态
     * 
     */
    public int? AuditFlag { get; set; }
    /**
     * 审核备注
     * 
     */
    public string? AuditRemark { get; set; }
    /**
     * 审核时间
     * 
     */
    public string? AuditTime { get; set; }
    /**
     * 分类ID
     * 
     */
    public string? CategoryId { get; set; }
    /**
     * 分类名称
     * 
     */
    public string? CategoryName { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? CategoryTag { get; set; }
    /**
     * 
     * 
     */
    public List<UserCouponEntity>? Coupons { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 备注
     * 
     */
    public string? Description { get; set; }
    /**
     * 启用
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? FakeTag { get; set; }
    /**
     * 
     * 
     */
    public long? FreeCount { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 封面图
     * 
     */
    public string? Image { get; set; }
    /**
     * 所属类型
     * 
     */
    public string? Mode { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 所属用户ID
     * 
     */
    public string? OwnerId { get; set; }
    /**
     * 单价
     * 
     */
    public int? Price { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? RecommendTag { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 
     * 
     */
    public long? TotalCount { get; set; }
    /**
     * 解锁方式
     * 
     */
    public string? UnlockType { get; set; }
    /**
     * 
     * 
     */
    public bool? Unlocked { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    /**
     * 
     * 
     */
    public long? VipCount { get; set; }
    
}




/**
 * 广告推广
 * 
 */
public class AdPromotionEntity
{
    /**
     * 开始时间
     * 
     */
    public string? BeginTime { get; set; }
    /**
     * 渠道
     * 
     */
    public string? Channel { get; set; }
    /**
     * 创建时间
     * 
     */
    public string? CreatedAt { get; set; }
    /**
     * 备注
     * 
     */
    public string? Description { get; set; }
    /**
     * 启用标识
     * 
     */
    public int? EnableFlag { get; set; }
    /**
     * 结束时间
     * 
     */
    public string? EndTime { get; set; }
    /**
     * 渠道应用icon
     * 
     */
    public string? Icon { get; set; }
    /**
     * 主键
     * 
     */
    public long? Id { get; set; }
    /**
     * 图片
     * 
     */
    public string? Image { get; set; }
    /**
     * 任务图片为主
     * 
     */
    public int? ImagePosition { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Installs { get; set; }
    /**
     * 跳转地址
     * 
     */
    public string? JumpUrl { get; set; }
    /**
     * 唤起标识
     * 
     */
    public int? LaunchFlag { get; set; }
    /**
     * 唤起地址
     * 
     */
    public string? LaunchUrl { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? MistakeAnswers { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Rewards { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? RightAnswers { get; set; }
    /**
     * 风控时间(秒)
     * 
     */
    public int? RiskInterval { get; set; }
    /**
     * 排序
     * 
     */
    public int? Sort { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Tips { get; set; }
    /**
     * 解锁任务阈值
     * 
     */
    public int? UnlockThreshold { get; set; }
    /**
     * 解锁类型
     * 
     */
    public int? UnlockType { get; set; }
    /**
     * 更新时间
     * 
     */
    public string? UpdatedAt { get; set; }
    /**
     * 视频教程
     * 
     */
    public string? Video { get; set; }
    
}




/**
 * 兴趣社
 * 
 */
public class AdInterestResult
{
    /**
     * 语音包收入
     * 
     */
    public List<UserItem>? AudioIncome { get; set; }
    /**
     * 本月冠军
     * 
     */
    public List<UserItem>? Champion { get; set; }
    /**
     * 合作伙伴
     * 
     */
    public List<UserItem>? Partner { get; set; }
    /**
     * 语音包排行
     * 
     */
    public List<UserItem>? Ranking { get; set; }
    /**
     * 视频收入
     * 
     */
    public List<UserItem>? VideoIncome { get; set; }
    /**
     * 热门视频
     * 
     */
    public List<VideoItem>? Videos { get; set; }
    
}




